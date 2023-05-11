package cdkemrserverlesswithdeltalake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-emrserverless-with-delta-lake-go/cdkemrserverlesswithdeltalake/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-emrserverless-with-delta-lake-go/cdkemrserverlesswithdeltalake/v2/internal"
)

// Creates a Service Catalog for EMR cluster templates.
//
// For detail, please refer to [Create AWS CloudFormation templates for Amazon EMR Studio](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-studio-cluster-templates.html).
//
// ```ts
// const emrClusterTemplatePortfolio = new EmrStudioDeveloperStack(this, 'ClusterTempalte');
// ```.
type EmrStudioDeveloperStack interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// The representative of the service catalog for EMR cluster tempaltes.
	Portfolio() awsservicecatalog.Portfolio
	// The representative of the product for demo purpose.
	Product() awsservicecatalog.Product
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for EmrStudioDeveloperStack
type jsiiProxy_EmrStudioDeveloperStack struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EmrStudioDeveloperStack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioDeveloperStack) Portfolio() awsservicecatalog.Portfolio {
	var returns awsservicecatalog.Portfolio
	_jsii_.Get(
		j,
		"portfolio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioDeveloperStack) Product() awsservicecatalog.Product {
	var returns awsservicecatalog.Product
	_jsii_.Get(
		j,
		"product",
		&returns,
	)
	return returns
}


func NewEmrStudioDeveloperStack(scope constructs.Construct, name *string, props *EmrStudioDeveloperStackProps) EmrStudioDeveloperStack {
	_init_.Initialize()

	if err := validateNewEmrStudioDeveloperStackParameters(scope, name, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrStudioDeveloperStack{}

	_jsii_.Create(
		"cdk-emrserverless-with-delta-lake.EmrStudioDeveloperStack",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

func NewEmrStudioDeveloperStack_Override(e EmrStudioDeveloperStack, scope constructs.Construct, name *string, props *EmrStudioDeveloperStackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-emrserverless-with-delta-lake.EmrStudioDeveloperStack",
		[]interface{}{scope, name, props},
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
func EmrStudioDeveloperStack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmrStudioDeveloperStack_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-emrserverless-with-delta-lake.EmrStudioDeveloperStack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudioDeveloperStack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

