package aws

import (
	"fmt"
	"mime/multipart"
)

type AwsHandler struct{}

func (a AwsHandler) Handle(file *multipart.FileHeader) {
	fmt.Println("handling s3 resource in aws environment")
}
