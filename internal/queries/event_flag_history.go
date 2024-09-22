package queries

import _ "embed"

//go:embed event_flag_history/insert.sql
var InsertFlagHistory string

//go:embed event_flag_history/select-by-redeemer.sql
var SelectFlagHistoryByRedeemer string
