package cdkemrserverlesswithdeltalake

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for the EMR Serverless bucket.
type EmrServerlessBucketProps struct {
	// The bucket name for EMR Serverless applications.
	// Default: - 'emr-serverless-AWS::AccountId'.
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Policy to apply when the bucket is removed from this stack.
	// Default: - The bucket will be deleted.
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

