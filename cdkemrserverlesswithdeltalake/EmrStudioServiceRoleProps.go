package cdkemrserverlesswithdeltalake


// Properties for defining the [service role](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-studio-service-role.html) of an EMR Studio.
type EmrStudioServiceRoleProps struct {
	// The custom construct as the workspace S3 bucket.
	WorkSpaceBucket WorkSpaceBucket `field:"required" json:"workSpaceBucket" yaml:"workSpaceBucket"`
	// A name for the service role of an EMR Studio.
	//
	// For valid values, see the RoleName parameter for the CreateRole action in the IAM API Reference.
	//
	// IMPORTANT: If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	//
	// If you specify a name, you must specify the CAPABILITY_NAMED_IAM value to acknowledge your template's capabilities. For more information, see Acknowledging IAM Resources in AWS CloudFormation Templates.
	// Default: - 'emr-studio-service-role'.
	//
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
}

