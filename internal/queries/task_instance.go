package queries

import _ "embed"

//go:embed task_instance/insert.sql
var InsertTaskInstance string

//go:embed task_instance/select.sql
var SelectTaskInstance string

//go:embed task_instance/update.sql
var UpdateTaskInstance string

//go:embed task_instance/delete.sql
var DeleteTaskInstance string
