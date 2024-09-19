package vpc_security_rule

type State string

const (
	Authorized State = "authorized"
	Revoked          = "revoked"
	Destroyed        = "destroyed"
)
