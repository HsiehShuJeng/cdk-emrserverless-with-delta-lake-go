package cdkemrserverlesswithdeltalake

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Interface for engine security group of EMR Studio.
type EmrStudioEngineSecurityGroupProps struct {
	// The VPC in which to create the engine security group for EMR Studio.
	// Default: - default VPC in an AWS account.
	//
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
}

