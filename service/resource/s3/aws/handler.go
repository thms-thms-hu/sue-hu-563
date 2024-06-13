package aws

import (
	"mime/multipart"
	"sue-hu-563/service/client/command"
	"sue-hu-563/util/yaml"
)

type AwsHandler struct{}

func (a AwsHandler) Handle(file *multipart.FileHeader) {
	content := yaml.ExtractContent(file)
	command.ApplyResource(content)
}
