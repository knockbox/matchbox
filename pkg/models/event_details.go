package models

// EventDetails represents the details for an Event
type EventDetails struct {
	Id             uint   `db:"id"`
	EventId        uint   `db:"event_id"`
	ProfilePicture string `db:"profile_picture"`
	Description    string `db:"description"`
	GithubURL      string `db:"github_url"`
	TwitterURL     string `db:"twitter_url"`
	WebsiteURL     string `db:"website_url"`
}
