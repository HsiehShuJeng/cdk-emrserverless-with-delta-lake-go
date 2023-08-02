package cdkemrserverlesswithdeltalake


// Options for the EMR Studio, mainly for EMR Serverless applications.
type EmrStudioProps struct {
	// The custom construct as the workspace S3 bucket.
	WorkSpaceBucket WorkSpaceBucket `field:"required" json:"workSpaceBucket" yaml:"workSpaceBucket"`
	// Specifies whether the Studio authenticates users using AWS SSO or IAM.
	// Default: - StudioAuthMode.AWS_IAM.
	//
	AuthMode StudioAuthMode `field:"optional" json:"authMode" yaml:"authMode"`
	// A detailed description of the Amazon EMR Studio.
	// Default: - 'EMR Studio Quick Launch - by scott.hsieh'
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the Amazon EMR Studio Engine security group.
	//
	// The Engine security group allows inbound network traffic from the Workspace security group, and it must be in the same VPC specified by VpcId.
	// Default: - a security group created by `EmrStudioEngineSecurityGroup`.
	//
	EngineSecurityGroupId *string `field:"optional" json:"engineSecurityGroupId" yaml:"engineSecurityGroupId"`
	// Options for which kind of identity will be associated with the Product of the Porfolio in AWS Service Catalog for EMR cluster templates.
	//
	// You can choose either an IAM group, IAM role, or IAM user. If you leave it empty, an IAM user named `Administrator` with the `AdministratorAccess` power needs to be created first.
	ServiceCatalogProps *EmrStudioDeveloperStackProps `field:"optional" json:"serviceCatalogProps" yaml:"serviceCatalogProps"`
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// A name for the service role of an EMR Studio.
	//
	// For valid values, see the RoleName parameter for the CreateRole action in the IAM API Reference.
	//
	// IMPORTANT: If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	//
	// If you specify a name, you must specify the CAPABILITY_NAMED_IAM value to acknowledge your template's capabilities. For more information, see Acknowledging IAM Resources in AWS CloudFormation Templates.
	// Default: - 'emr-studio-service-role'.
	//
	ServiceRoleName *string `field:"optional" json:"serviceRoleName" yaml:"serviceRoleName"`
	// A descriptive name for the Amazon EMR Studio.
	// Default: - 'emr-sutdio-quicklaunch'.
	//
	StudioName *string `field:"optional" json:"studioName" yaml:"studioName"`
	// The subnet IDs for the EMR studio.
	//
	// You can select the subnets from the default VPC in your AWS account.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The custom user role for the EMR Studio when authentication is AWS SSO.
	//
	// Currently, if you choose to establish an EMR serverless application where the authentication mechanism used by the EMR Studio is AWS SSO, you need to create a user role by yourself and assign the role arn to this argument if AWS SSO is chosen as authentication for the EMR Studio;.
	UserRoleArn *string `field:"optional" json:"userRoleArn" yaml:"userRoleArn"`
	// Used by the EMR Studio.
	// Default: - 'The default VPC will be used.'
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// The ID of the security group used by the workspace.
	// Default: - a security group created by `EmrStudioWorkspaceSecurityGroup`.
	//
	WorkSpaceSecurityGroupId *string `field:"optional" json:"workSpaceSecurityGroupId" yaml:"workSpaceSecurityGroupId"`
}

