package queries

import _ "embed"

//go:embed deployment/insert.sql
var InsertDeployment string

//go:embed deployment/select-by-event_id.sql
var SelectDeploymentByEventId string

//go:embed deployment/update-status.sql
var UpdateDeploymentStatusById string
