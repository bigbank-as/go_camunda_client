package dto

type Process struct {
	JsonId      string `json:"id"`
	JsonIsEnded bool   `json:"ended"`
}

func (process Process) GetId() string {
	return process.JsonId
}

func (process Process) IsEnded() bool {
	return process.JsonIsEnded
}
