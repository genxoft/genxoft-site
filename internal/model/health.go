package model

type Health struct {
	Status    string `json:"message"`
	Version   string `json:"version"`
	ReleaseId string `json:"releaseID"`
}

func NewHealth(message string, version string, releaseId string) *Health {
	return &Health{
		Status:    message,
		Version:   version,
		ReleaseId: releaseId,
	}
}
