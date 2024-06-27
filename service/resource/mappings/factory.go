package mappings

import (
	"sue-hu-563/service/resource/mappings/azure/aws"
)

func GetFieldMapping(environmentFrom string, environmentTo string, resource string) map[string]string {
	switch environmentTo {
	case "azure":
		switch environmentFrom {
		case "aws":
			switch resource {
			case "s3":
				return aws.AWSAzureFieldMapping
			}

		}
	}
	panic("no mapping found!")
}

func GetValueMapping(environmentFrom string, environmentTo string, resource string) map[string]string {
	switch environmentTo {
	case "azure":
		switch environmentFrom {
		case "aws":
			switch resource {
			case "s3":
				return aws.AWSAzureValueMapping
			}

		}
	}
	panic("no mapping found!")
}

func GetRegionMapping(environmentFrom string, environmentTo string) map[string]string {
	switch environmentTo {
	case "azure":
		switch environmentFrom {
		case "aws":
			return aws.AWSAzureRegionMapping
		}
	}
	panic("no mapping found!")
}
