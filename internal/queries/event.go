package queries

import _ "embed"

//go:embed event/insert.sql
var InsertEvent string

//go:embed event/select-all.sql
var SelectAllEvents string

//go:embed event/select-by-activity_id.sql
var SelectEventByActivityId string
