package s3

import "mime/multipart"

type S3Handler interface {
	Handle(file []*multipart.FileHeader)
}
