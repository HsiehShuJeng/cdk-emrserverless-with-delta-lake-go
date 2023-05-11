package cdkemrserverlesswithdeltalake

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Options for the execution job role of EMR Serverless.
type ServerlessJobRoleProps struct {
	// The EMR Serverless bucket.
	EmrServerlessBucket awss3.Bucket `field:"required" json:"emrServerlessBucket" yaml:"emrServerlessBucket"`
}

