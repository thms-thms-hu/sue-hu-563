package aws

import (
	"mime/multipart"
)

type AwsHandler struct{}

func (a AwsHandler) Handle(file *multipart.FileHeader) {
	// TODO: VALIDATION LOGIC

	// TODO: APPLY RESOURCE
}
