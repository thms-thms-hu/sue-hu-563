package service

import "mime/multipart"

func ApplyResource(resource string, environment string, file *multipart.FileHeader) {
	switch resource {
	case "s3":
		// todo: factory call -> factory result function call
	}
}
