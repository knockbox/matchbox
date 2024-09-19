package docker

// CheckRepositoryTagOptions contains the fields required to make a request.
type CheckRepositoryTagOptions struct {
	Namespace  string
	Repository string
	Tag        string
}
