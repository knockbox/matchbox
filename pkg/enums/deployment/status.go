package deployment

type Status string

const (
	Preparing Status = "preparing"
	Idle             = "idle"
	Ready            = "ready"
	Live             = "live"
	Teardown         = "teardown"
	Complete         = "complete"
)
