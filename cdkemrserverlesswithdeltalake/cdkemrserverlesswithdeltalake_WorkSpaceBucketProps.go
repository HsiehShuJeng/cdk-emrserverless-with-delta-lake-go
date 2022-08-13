// A construct for the quick demo of EMR Serverless.
package cdkemrserverlesswithdeltalake

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

type WorkSpaceBucketProps struct {
	// The bucket name for the workspace of an EMR Studio.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Policy to apply when the bucket is removed from this stack.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

