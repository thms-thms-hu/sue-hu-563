package aws

import (
	"mime/multipart"
	"sue-hu-563/service/client/command"
	"sue-hu-563/util/yaml"
)

type AwsHandler struct{}

func (a AwsHandler) Handle(files []*multipart.FileHeader) {
	// TODO: VALIDATION LOGIC (afhankelijkheden zoals providerconfig of bucketversioning die afhankelijk is van een bucket)

	// TODO:
	//  unique names? naam van {klant}-{guid}
	//  secrets + providerconfig. bij opzetten van cloud infra of nog apart aanbrengen/controleren?
	//  rollback na error?
	//  meerdere bestanden + --- divider?

	// TODO: de volgorde moet misschien nog goed verlopen onafhankelijk van volgorde uploaden in request

	for _, file := range files {
		content := yaml.ExtractContent(file)
		command.ApplyResource(content)
	}
}
