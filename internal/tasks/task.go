package tasks

import (
	"bulk-file-converter/internal/jobs"
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/hibiken/asynq"
)

const (
	TypeFileConvert = "file:convert"
)

type FileConvertPayload struct {
	Bucket    string `json:"bucket"`
	ObjectKey string `json:"object_key"`
	TargetExt string `json:"target_ext"`
}

func NewFileConvertTask(p FileConvertPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(p)
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
	sourceExt := filepath.Ext(p.ObjectKey)
	if len(sourceExt) > 0 {
		sourceExt = sourceExt[1:]
	}

	switch {
	case sourceExt == "pdf" && p.TargetExt == "jpg":
		jobs.ConvertPDFToImage(p)
	case sourceExt == "pdf" && p.TargetExt == "txt":
		jobs.ConvertPDFToText(p)

	case sourceExt == "jpg" && p.TargetExt == "pdf":
		jobs.ConvertImageToPDF(p)

	default:
		return fmt.Errorf("unsupported conversion: %s to %s", sourceExt, p.TargetExt)
	}

	return nil
}
