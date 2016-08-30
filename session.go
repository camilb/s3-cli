package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const DEFAULT_REGION = "us-east-1"

func SessionNew(config *Config) *s3.S3 {
	//	creds := credentials.NewStaticCredentials(config.AccessKey, config.SecretKey, "")

	// By default make sure a region is specified Add  "Credentials: creds" to enable access with credentials
	return s3.New(session.New(&aws.Config{Region: aws.String(DEFAULT_REGION)}))
}

func SessionForBucket(svc *s3.S3, bucket string) (*s3.S3, error) {
	params := &s3.HeadBucketInput{Bucket: aws.String(bucket)}
	_, err := svc.HeadBucket(params)
	if err != nil {
		return nil, err
	}

	if loc, err := svc.GetBucketLocation(&s3.GetBucketLocationInput{Bucket: &bucket}); err != nil {
		return nil, err
	} else if loc.LocationConstraint != nil {
		return s3.New(session.New(&svc.Client.Config, &aws.Config{Region: loc.LocationConstraint})), nil
	}
	return svc, nil
}
