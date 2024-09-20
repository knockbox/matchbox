package queries

import _ "embed"

//go:embed event_details/insert.sql
var InsertEventDetails string

//go:embed event_details/update.sql
var UpdateEventDetails string
