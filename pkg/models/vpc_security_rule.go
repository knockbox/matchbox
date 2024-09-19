package models

import "github.com/knockbox/matchbox/pkg/enums/vpc_security_rule"

// VPCSecurityRule represents a security rule for a given VPC
type VPCSecurityRule struct {
	Id            uint                    `db:"id"`
	VPCInstanceId uint                    `db:"vpc_instance_id"`
	IPAddress     string                  `db:"ip_address"`
	Ingress       bool                    `db:"ingress"`
	Egress        bool                    `db:"egress"`
	State         vpc_security_rule.State `db:"state"`
}
