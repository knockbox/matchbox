package payloads

import (
	"github.com/knockbox/matchbox/pkg/enums/event"
)

type EventParticipantCreate struct {
	Status    event.Status `json:"status" validate:"required"`
	CanInvite *bool        `json:"can_invite" validate:"required,boolean"`
	CanManage *bool        `json:"can_manage" validate:"required,boolean"`
}
