package queries

import _ "embed"

//go:embed event_participant/insert.sql
var InsertParticipant string

//go:embed event_participant/update.sql
var UpdateParticipant string

//go:embed event_participant/select-all.sql
var SelectAllParticipants string
