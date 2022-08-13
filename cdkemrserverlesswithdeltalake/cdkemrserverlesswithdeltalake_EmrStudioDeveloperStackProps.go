// A construct for the quick demo of EMR Serverless.
package cdkemrserverlesswithdeltalake

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Interface for Service Catalog of EMR cluster templates.
type EmrStudioDeveloperStackProps struct {
	// an IAM group you wish to associate with the Portfolio for EMR cluster template.
	Group awsiam.IGroup `field:"optional" json:"group" yaml:"group"`
	// The provider name in a Service Catalog for EMR cluster templates.
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
	// an IAM role you wish to associate with the Portfolio for EMR cluster template.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// an IAM user you wish to associate with the Portfolio for EMR cluster template.
	User awsiam.IUser `field:"optional" json:"user" yaml:"user"`
}

