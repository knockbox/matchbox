package enums

type EventStatus string

const (
	Invited   EventStatus = "invited"
	Declined              = "declined"
	Member                = "member"
	Requested             = "requested"
	Removed               = "removed"
	Banned                = "banned"
)
