package aws

var AWSAzureFieldMapping = map[string]string{
	"apiVersion":              "apiVersion",
	"kind":                    "kind",
	"metadata.name":           "metadata.name",
	"spec.forProvider.region": "spec.forProvider.location",
}

var AWSAzureValueMapping = map[string]string{
	"Bucket": "BlobStorage",
}
