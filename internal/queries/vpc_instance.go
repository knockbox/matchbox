package queries

import _ "embed"

//go:embed vpc_instance/insert.sql
var InsertVPCInstance string

//go:embed vpc_instance/select.sql
var SelectVPCInstance string
