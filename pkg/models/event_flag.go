package models

import (
	"github.com/google/uuid"
	"github.com/knockbox/matchbox/pkg/enums"
)

// EventFlag represents a generated for an Event.
type EventFlag struct {
	Id         uint             `db:"id"`
	EventId    uint             `db:"event_id"`
	FlagId     uuid.UUID        `db:"flag_id"`
	Difficulty enums.Difficulty `db:"difficulty"`
	EnvVar     string           `db:"env_var"`
}
