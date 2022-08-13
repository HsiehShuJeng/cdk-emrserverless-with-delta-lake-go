// A construct for the quick demo of EMR Serverless.
package cdkemrserverlesswithdeltalake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-emrserverless-with-delta-lake-go/cdkemrserverlesswithdeltalake/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-emrserverless-with-delta-lake-go/cdkemrserverlesswithdeltalake/v2/internal"
)

// Creates a Lambda function for the custom resource which can add necessary tag onto the VPC and subnets for the EMR Studio during deployment.
//
// For detail on the tag, please refer to [How to create a service role for EMR Studio](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-studio-service-role.html#emr-studio-service-role-instructions)
type EmrStudioTaggingExpert interface {
	constructs.Construct
	// The repesentative of the Lambda function for the custom resource which can add necessary tag onto the VPC and subnets for the EMR Studio during deployment.
	FunctionEntity() awslambda.Function
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for EmrStudioTaggingExpert
type jsiiProxy_EmrStudioTaggingExpert struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EmrStudioTaggingExpert) FunctionEntity() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"functionEntity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioTaggingExpert) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewEmrStudioTaggingExpert(scope constructs.Construct, name *string) EmrStudioTaggingExpert {
	_init_.Initialize()

	j := jsiiProxy_EmrStudioTaggingExpert{}

	_jsii_.Create(
		"cdk-emrserverless-with-delta-lake.EmrStudioTaggingExpert",
		[]interface{}{scope, name},
		&j,
	)

	return &j
}

func NewEmrStudioTaggingExpert_Override(e EmrStudioTaggingExpert, scope constructs.Construct, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-emrserverless-with-delta-lake.EmrStudioTaggingExpert",
		[]interface{}{scope, name},
		e,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func EmrStudioTaggingExpert_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-emrserverless-with-delta-lake.EmrStudioTaggingExpert",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudioTaggingExpert) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

