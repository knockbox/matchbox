package payloads

type EventDetailsUpdate struct {
	ProfilePicture *string `json:"profile_picture,omitempty" validator:"omitempty,http_url"`
	Description    *string `json:"description,omitempty" validator:"omitempty,lte=4096"`
	GithubURL      *string `json:"github_url,omitempty" validator:"omitempty,http_url"`
	TwitterURL     *string `json:"twitter_url,omitempty" validator:"omitempty,http_url"`
	WebsiteURL     *string `json:"website_url,omitempty" validator:"omitempty,http_url"`
}
