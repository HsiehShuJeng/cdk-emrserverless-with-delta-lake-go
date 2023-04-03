// A construct for the quick demo of EMR Serverless.
package cdkemrserverlesswithdeltalake

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Interface for workspace security group of EMR Studio.
type EmrStudioWorkspaceSecurityGroupProps struct {
	// The VPC in which to create workspace security group for EMR Studio.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
}

