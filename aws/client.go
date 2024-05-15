package aws

import (
	"github.com/tappoy/archive/types"

	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	client *s3.Client
	bucket string
}

// NewClient is a factory method for S3Client
func NewClient(region, accessKey, secretKey, endpoint, bucket string) (types.Client, error) {
	appCreds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""))

	cfg := aws.Config{
		Region:       region,
		Credentials:  appCreds,
		BaseEndpoint: &endpoint,
	}

	return S3Client{client: s3.NewFromConfig(cfg), bucket: bucket}, nil
}

// NewClientFromConfig is a factory method for OpenstackClient.
func NewClientFromConfig(config map[string]string) (types.Client, error) {
	// check required fields
	if _, ok := config["AWS_REGION"]; !ok {
		return nil, fmt.Errorf("missing AWS_REGION")
	}

	if _, ok := config["AWS_ACCESS_KEY_ID"]; !ok {
		return nil, fmt.Errorf("missing AWS_ACCESS_KEY_ID")
	}

	if _, ok := config["AWS_SECRET_ACCESS_KEY"]; !ok {
		return nil, fmt.Errorf("missing AWS_SECRET_ACCESS_KEY")
	}

	if _, ok := config["AWS_ENDPOINT"]; !ok {
		return nil, fmt.Errorf("missing AWS_ENDPOINT")
	}

	if _, ok := config["AWS_BUCKET"]; !ok {
		return nil, fmt.Errorf("missing AWS_BUCKET")
	}

	return NewClient(config["AWS_REGION"], config["AWS_ACCESS_KEY_ID"], config["AWS_SECRET_ACCESS_KEY"], config["AWS_ENDPOINT"], config["AWS_BUCKET"])
}
