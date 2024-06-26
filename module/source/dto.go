package source

type SourceDTO struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	ProjectId   int64  `json:"projectId"`
}

type SourceDetail struct {
	*SourceDTO
	Config string `json:"config"`
}
