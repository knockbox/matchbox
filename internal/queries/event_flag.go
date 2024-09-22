package queries

import _ "embed"

//go:embed event_flag/insert.sql
var InsertEventFlag string

//go:embed event_flag/update.sql
var UpdateEventFlag string

//go:embed event_flag/select-by-flag_id.sql
var SelectEventFlagByFlagId string

//go:embed event_flag/select-all.sql
var SelectAllEventFlags string

//go:embed event_flag/delete.sql
var DeleteEventFlag string
