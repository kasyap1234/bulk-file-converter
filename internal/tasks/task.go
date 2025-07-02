package tasks

import (
	"bulk-file-converter/internal/jobs"
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

const (
	TypeFileConvert = "file:convert"
)

type FileConvertPayload struct {
	PayloadID int    `json:"payload_id"`
	FileName  string `json:"filename"`
	FileExt   string `json:"fileExt"`
	TargetExt string `json:"targetExt"`
}

func NewFileConvertTask(fileconvertpayload FileConvertPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(fileconvertpayload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeFileConvert, payload), nil
}

func HandleFileConvertTask(ctx context.Context, t *asynq.Task) error {
	var p FileConvertPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed %v %w", err, asynq.SkipRetry)
	}
	switch {
	case p.FileExt == "pdf" && p.TargetExt == "jpg":
		jobs.ConvertPDFToImage(p)
	case p.FileExt == "pdf" && p.TargetExt == "txt":
		jobs.ConvertPDFToText(p)

	case p.FileExt == "jpg" && p.TargetExt == "pdf":
		jobs.ConvertImageToPDF(p)

	default:
		return fmt.Errorf("unsupported conversion: %s to %s", p.FileExt, p.TargetExt)
	}

	return nil
}
