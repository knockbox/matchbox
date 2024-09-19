package payloads

type EventCreate struct {
	Name            string `json:"name" validate:"required,gte=1,lte=64"`
	StartsAt        int64  `json:"starts_at" validate:"required"`
	EndsAt          int64  `json:"ends_at" validate:"required,gtcsfield=StartsAt"`
	ImageNamespace  string `json:"image_namespace" validate:"required,gte=1,lte=256"`
	ImageRepository string `json:"image_repository" validate:"required,gte=1,lte=256"`
	ImageTag        string `json:"image_tag" validate:"required"`
	Private         *bool  `json:"private" validate:"required,boolean"`
}
