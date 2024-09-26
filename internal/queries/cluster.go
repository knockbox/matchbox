package queries

import _ "embed"

//go:embed cluster/insert.sql
var InsertCluster string

//go:embed cluster/select.sql
var SelectCluster string
