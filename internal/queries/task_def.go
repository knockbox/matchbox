package queries

import _ "embed"

//go:embed task_def/insert.sql
var InsertTaskDef string

//go:embed task_def/select.sql
var SelectTaskDefByDeploymentId string
