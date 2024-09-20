package models

import "github.com/knockbox/matchbox/pkg/payloads"

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

func (d *EventDetails) ApplyUpdate(payload *payloads.EventDetailsUpdate) {
	if payload.ProfilePicture != nil {
		d.ProfilePicture = *payload.ProfilePicture
	}

	if payload.Description != nil {
		d.Description = *payload.Description
	}

	if payload.GithubURL != nil {
		d.GithubURL = *payload.Description
	}

	if payload.TwitterURL != nil {
		d.TwitterURL = *payload.TwitterURL
	}

	if payload.WebsiteURL != nil {
		d.WebsiteURL = *payload.WebsiteURL
	}
}
