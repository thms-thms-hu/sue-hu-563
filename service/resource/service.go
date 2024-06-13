package resource

import (
	"mime/multipart"
	"sue-hu-563/service/resource/s3"
)

func ApplyResource(resource string, environment string, file *multipart.FileHeader) {
	// 1. identify environment of desired resource yaml
	// 2. if identified env == chosen env -> step 6
	// 3. if identified env != chosen env
	// 4. extract values from keys of yaml file
	// 5. place them in yaml config that conforms to chosen env
	// 6. push file

	switch resource {
	case "s3":
		handler := s3.GetS3Handler(environment)
		handler.Handle(file)
	}
}
