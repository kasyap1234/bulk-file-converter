package tasks

const TypeConverterPDFTOImage = "converter:pdf_to_image"

type TypeConverterPDFToImagePayload struct {
	BucketName string `json:"bucket_name"`
	ObjectKey  string `json:"object_key"`
}
