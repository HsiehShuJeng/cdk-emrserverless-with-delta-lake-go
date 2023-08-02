package cdkemrserverlesswithdeltalake


type EmrServerlessProps struct {
	// Options for which kind of identity will be associated with the Product of the Porfolio in AWS Service Catalog for EMR cluster templates.
	//
	// You can choose either an IAM group, IAM role, or IAM user. If you leave it empty, an IAM user named `Administrator` with the `AdministratorAccess` power needs to be created first.
	ServiceCatalogProps *EmrStudioDeveloperStackProps `field:"optional" json:"serviceCatalogProps" yaml:"serviceCatalogProps"`
	// The subnet IDs for the EMR studio.
	//
	// You can select the subnets from the default VPC in your AWS account.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Used by the EMR Studio.
	// Default: - 'The default VPC will be used.'
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

