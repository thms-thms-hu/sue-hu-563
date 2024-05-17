package s3

import "sue-hu-563/service/resource/s3/aws"

func GetS3Handler(environment string) S3Handler {
	switch environment {
	case "aws":
		return aws.AwsHandler{}
	}

	// TODO: throw error
	return nil
}
