package event

type Status string

const (
	Invited   Status = "invited"
	Declined         = "declined"
	Member           = "member"
	Requested        = "requested"
	Removed          = "removed"
	Banned           = "banned"
)
