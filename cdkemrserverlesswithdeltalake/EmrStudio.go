package cdkemrserverlesswithdeltalake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-emrserverless-with-delta-lake-go/cdkemrserverlesswithdeltalake/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsemr"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-emrserverless-with-delta-lake-go/cdkemrserverlesswithdeltalake/v2/internal"
)

// Creates an EMR Studio for EMR Serverless applications.
//
// The Studio is not only for EMR Serverless applications but also for launching an EMR cluster via a cluster template created in this constrcut to check out results transformed by EMR serverless applications.
//
// For what Studio can do further, please refer to [Amazon EMR Studio](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-studio.html).
//
// ```ts
// const workspaceBucket = new WorkSpaceBucket(this, 'EmrStudio');
// const emrStudio = new EmrStudio(this, '', {
//    workSpaceBucket: workspaceBucket,
//    subnetIds: ['subnet1', 'subnet2', 'subnet3']
// });
// ```.
type EmrStudio interface {
	constructs.Construct
	Entity() awsemr.CfnStudio
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for EmrStudio
type jsiiProxy_EmrStudio struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EmrStudio) Entity() awsemr.CfnStudio {
	var returns awsemr.CfnStudio
	_jsii_.Get(
		j,
		"entity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewEmrStudio(scope constructs.Construct, name *string, props *EmrStudioProps) EmrStudio {
	_init_.Initialize()

	if err := validateNewEmrStudioParameters(scope, name, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrStudio{}

	_jsii_.Create(
		"cdk-emrserverless-with-delta-lake.EmrStudio",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

func NewEmrStudio_Override(e EmrStudio, scope constructs.Construct, name *string, props *EmrStudioProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-emrserverless-with-delta-lake.EmrStudio",
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
func EmrStudio_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmrStudio_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-emrserverless-with-delta-lake.EmrStudio",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudio) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

