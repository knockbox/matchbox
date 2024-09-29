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

	event        accessors.EventAccessor
	eventDetails accessors.EventDetailsAccessor
	flag         accessors.EventFlagAccessor
	participant  accessors.EventParticipantAccessor
	flagHistory  accessors.EventFlagHistoryAccessor

	l hclog.Logger
}

// NewEventClient creates a new EventClient using the SQLImpl accessors.
func NewEventClient(db *sqlx.DB, l hclog.Logger) *EventClient {
	return &EventClient{
		dc: docker.NewClient(l),
		event: platform.EventSQLImpl{
			DB: db,
		},
		eventDetails: platform.EventDetailsDQLImpl{
			DB: db,
		},
		flag: platform.EventFlagSQLImpl{
			DB: db,
		},
		participant: platform.EventParticipantSQLImpl{
			DB: db,
		},
		flagHistory: platform.EventFlagHistorySQLImpl{
			DB: db,
		},
		l: l,
	}
}

func (e *EventClient) CreateEvent(payload *payloads.EventCreate, organizer uuid.UUID) (*models.Event, error) {
	event := models.NewEvent(organizer)
	if err := event.ApplyCreate(payload); err != nil {
		return nil, err
	}

	dockerResult := e.dc.CheckRepositoryTag(context.Background(), &docker.CheckRepositoryTagOptions{
		Namespace:  event.ImageName,
		Repository: event.ImageRepo,
		Tag:        event.ImageTag,
	})
	if dockerResult.Error != nil {
		return nil, dockerResult.Error
	}
	if !dockerResult.Exists || dockerResult.Private {
		return nil, ErrDockerTagFailed
	}

	result, err := e.event.Create(*event)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	if _, err := e.eventDetails.CreateForEvent(int(id)); err != nil {
		return nil, err
	}

	return event, err
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

func (e *EventClient) UpdateEventDetails(details *models.EventDetails, payload *payloads.EventDetailsUpdate) error {
	details.ApplyUpdate(payload)
	_, err := e.eventDetails.Update(*details)
	return err
}

func (e *EventClient) CreateFlag(event *models.Event, payload *payloads.EventFlagCreate) error {
	flag := models.NewEventFlag(event.Id)
	flag.ApplyCreate(payload)

	_, err := e.flag.Create(*flag)
	return err
}

func (e *EventClient) UpdateFlag(event *models.Event, flagId uuid.UUID, payload *payloads.EventFlagUpdate) error {
	flag := models.NewEventFlag(event.Id)
	flag.ApplyUpdate(payload)
	flag.FlagId = flagId

	_, err := e.flag.Update(*flag)
	return err
}

func (e *EventClient) GetAllEventFlags(event *models.Event) ([]models.EventFlag, error) {
	return e.flag.GetAllForEvent(int(event.Id))
}

func (e *EventClient) GetEventFlagByFlagId(flagId uuid.UUID) (*models.EventFlag, error) {
	flag, err := e.flag.GetByFlagId(flagId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return flag, err
}

func (e *EventClient) DeleteEventFlag(flagId uuid.UUID) error {
	_, err := e.flag.DeleteByFlagId(flagId)
	return err
}

func (e *EventClient) CreateParticipant(event *models.Event, id uuid.UUID, payload *payloads.EventParticipantCreate) error {
	participant := models.NewEventParticipant(event, id)
	participant.ApplyCreate(payload)

	_, err := e.participant.Create(*participant)
	return err
}

func (e *EventClient) GetAllParticipants(event *models.Event) ([]models.EventParticipant, error) {
	return e.participant.GetAllByEventId(int(event.Id))
}

func (e *EventClient) GetParticipantByEventAndParticipantId(event *models.Event, participantId uuid.UUID) (*models.EventParticipant, error) {
	participant, err := e.participant.GetByEventAndParticipantId(int(event.Id), participantId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return participant, err
}

func (e *EventClient) RedeemFlag(event *models.Event, participant *models.EventParticipant, flag *models.EventFlag) error {
	history := models.NewFlagHistory(event, participant, flag)

	_, err := e.flagHistory.Create(*history)
	return err
}

func (e *EventClient) GetRedeemedFlag(event *models.Event, participant *models.EventParticipant, flag *models.EventFlag) (*models.EventFlagHistory, error) {
	history, err := e.flagHistory.GetByEventFlagRedeemer(int(event.Id), int(flag.Id), participant.ParticipantId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return history, err
}

func (e *EventClient) GetAllHistoryForEvent(event *models.Event) ([]models.EventFlagHistory, error) {
	return e.flagHistory.GetByEvent(int(event.Id))
}
