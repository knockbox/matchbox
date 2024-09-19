package docker

// CheckRepositoryTagResult contains the result for Client.CheckRepositoryTag
type CheckRepositoryTagResult struct {
	Exists  bool
	Private bool
	Error   error
}
