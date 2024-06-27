package resource

import (
	"k8s.io/utils/strings/slices"
	"mime/multipart"
	"sigs.k8s.io/yaml"
	"sue-hu-563/service/resource/mappings"
	"sue-hu-563/util/yaml_configs"
)

func ApplyResource(resource string, environment string, file *multipart.FileHeader) {
	// 1. identify environment of desired resource yaml_configs
	content := yaml_configs.ExtractContent(file)
	var result map[string]interface{}
	yaml.Unmarshal(content, &result)

	fileEnvironment := identifyEnvironment(result)

	if environment != fileEnvironment {
		fieldMapping := mappings.GetFieldMapping(fileEnvironment, environment, resource)
		valueMapping := mappings.GetValueMapping(fileEnvironment, environment, resource)
		regionMapping := mappings.GetRegionMapping(fileEnvironment, environment)
		convert(result, fieldMapping, valueMapping, regionMapping)
	}

	// push yaml

}

func identifyEnvironment(file map[string]interface{}) string {

	if slices.Contains(RegionAWS, yaml_configs.FindValueByKey("region", file)) {
		return "aws"
	}

	if slices.Contains(RegionAzure, yaml_configs.FindValueByKey("location", file)) {
		return "azure"
	}

	panic("could not identify environment based on provided file")
}

var RegionAWS = []string{"eu-central-1", "eu-west-1", "eu-west-2", "eu-south-1", "eu-west-3", "eu-south-2", "eu-north-1", "eu-central-2"}
var RegionAzure = []string{"europe", "norteurope", "westeurope", "francecentral", "germanywestcentral", "italynorth", "norwayeast", "polandcentral", "switzerlandnorth", "francesouth", "germanynorth", "norwaywest", "switzerlandwest", "ukwest"}
var RegionGCP = []string{"europe-west1", "europe-west2", "europe-west3", "europe-west4", "europe-west8", "europe-west9", "europe-west10", "europe-west12", "europe-southwest1", "europe-north1", "europe-central2"}
