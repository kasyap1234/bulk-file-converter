package models

type Status string

const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
	StatusFailed     Status = "failed"
)

type FileConversion struct {
	ID            int    `json:"id"`
	FileName      string `json:"fileName"`
	FileExt       string `json:"fileExt"`
	ConvertedName string `json:"convertedName"`
	ConvertedExt  string `json:"convertedExt"`
	Status        string `json:"status"`
}
