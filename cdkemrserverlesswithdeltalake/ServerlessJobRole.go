package cdkemrserverlesswithdeltalake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-emrserverless-with-delta-lake-go/cdkemrserverlesswithdeltalake/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-emrserverless-with-delta-lake-go/cdkemrserverlesswithdeltalake/v2/internal"
)

// Creates an execution job role for EMR Serverless.
//
// For detail, please refer to [Create a job runtime role](https://docs.aws.amazon.com/emr/latest/EMR-Serverless-UserGuide/getting-started.html#gs-runtime-role).
//
// ```ts
// const emrServerlessBucket = new EmrServerlessBucket(this, 'EmrServerlessStorage');
// const emrServerlessJobRole = new ServerlessJobRole(this, 'EmrServerlessJob', {emrServerlessBucket: emrServerlessBucket});
// ```.
type ServerlessJobRole interface {
	constructs.Construct
	// The representative of the execution role for EMR Serverless.
	Entity() awsiam.Role
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ServerlessJobRole
type jsiiProxy_ServerlessJobRole struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ServerlessJobRole) Entity() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"entity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessJobRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewServerlessJobRole(scope constructs.Construct, name *string, props *ServerlessJobRoleProps) ServerlessJobRole {
	_init_.Initialize()

	if err := validateNewServerlessJobRoleParameters(scope, name, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServerlessJobRole{}

	_jsii_.Create(
		"cdk-emrserverless-with-delta-lake.ServerlessJobRole",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

func NewServerlessJobRole_Override(s ServerlessJobRole, scope constructs.Construct, name *string, props *ServerlessJobRoleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-emrserverless-with-delta-lake.ServerlessJobRole",
		[]interface{}{scope, name, props},
		s,
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
func ServerlessJobRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServerlessJobRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-emrserverless-with-delta-lake.ServerlessJobRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessJobRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

