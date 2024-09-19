package client

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	"github.com/knockbox/matchbox/internal/platform"
	"github.com/knockbox/matchbox/pkg/accessors"
	"github.com/knockbox/matchbox/pkg/docker"
	"github.com/knockbox/matchbox/pkg/models"
	"github.com/knockbox/matchbox/pkg/payloads"
)

type EventClient struct {
	dc *docker.Client

	event accessors.EventAccessor
	l     hclog.Logger
}

// NewEventClient creates a new EventClient using the SQLImpl accessors.
func NewEventClient(db *sqlx.DB, l hclog.Logger) *EventClient {
	return &EventClient{
		dc: docker.NewClient(l),
		event: platform.EventSQLImpl{
			DB:     db,
			Logger: l,
		},
		l: l,
	}
}

func (e *EventClient) CreateEvent(payload *payloads.EventCreate, organizer uuid.UUID) error {
	event := models.NewEvent(organizer)
	if err := event.ApplyCreate(payload); err != nil {
		return err
	}

	dockerResult := e.dc.CheckRepositoryTag(context.Background(), &docker.CheckRepositoryTagOptions{
		Namespace:  event.ImageName,
		Repository: event.ImageRepo,
		Tag:        event.ImageTag,
	})
	if dockerResult.Error != nil {
		return dockerResult.Error
	}
	if !dockerResult.Exists || dockerResult.Private {
		return ErrDockerTagFailed
	}

	result, err := e.event.Create(*event)
	if err != nil {
		return err
	}

	// TODO: Create the Event Details
	e.l.Info("EventClient.CreateEvent", "result", result)

	return err
}

func (e *EventClient) GetAllEvents() ([]models.Event, error) {
	return e.event.GetAll()
}

func (e *EventClient) GetByActivityId(activityId string) (*models.Event, error) {
	event, err := e.event.GetByActivityId(activityId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return event, err
}
