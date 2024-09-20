package payloads

import (
	"github.com/knockbox/matchbox/pkg/enums/difficulty"
)

type EventFlagCreate struct {
	Difficulty difficulty.Difficulty `json:"difficulty" validate:"required"`
	EnvVar     string                `json:"env_var" validate:"required,gt=0,lte=128"`
}

type EventFlagUpdate struct {
	Difficulty *difficulty.Difficulty `json:"difficulty,omitempty" validate:"omitempty"`
	EnvVar     *string                `json:"env_var,omitempty" validate:"omitempty,gt=0,lte=128"`
}
