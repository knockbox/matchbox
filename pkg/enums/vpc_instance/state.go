package vpc_instance

type State string

const (
	Destroyed State = "destroyed"
	Pending         = "pending"
	Available       = "available"
)
