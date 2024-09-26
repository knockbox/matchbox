package queries

import _ "embed"

//go:embed efs/insert.sql
var InsertEFS string

//go:embed efs/select.sql
var SelectEFS string
