package resource

import (
	"mime/multipart"
	"sue-hu-563/service/resource/s3"
)

func ApplyResource(resource string, environment string, file []*multipart.FileHeader) {
	switch resource {
	case "s3":
		handler := s3.GetS3Handler(environment)
		handler.Handle(file)
	}
}
