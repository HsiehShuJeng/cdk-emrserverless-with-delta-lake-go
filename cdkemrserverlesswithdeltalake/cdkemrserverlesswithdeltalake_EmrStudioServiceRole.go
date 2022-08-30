// A construct for the quick demo of EMR Serverless.
package cdkemrserverlesswithdeltalake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-emrserverless-with-delta-lake-go/cdkemrserverlesswithdeltalake/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-emrserverless-with-delta-lake-go/cdkemrserverlesswithdeltalake/v2/internal"
)

// Creates a default service role for an EMR Studio.
//
// For detail, please refer to [Create an EMR Studio service role](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-studio-service-role.html).
//
// ```ts
// const workSpaceBucket = new WorkSpaceBucket(this, 'WorkSpace');
// const emrStudioServiceRole = new EmrStudioServiceRole(this, 'Service', {
//       workSpaceBucket: workSpaceBucket
// });
// ```.
type EmrStudioServiceRole interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// The representative of the default service role for EMR Studio.
	RoleEntity() awsiam.Role
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for EmrStudioServiceRole
type jsiiProxy_EmrStudioServiceRole struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EmrStudioServiceRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioServiceRole) RoleEntity() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"roleEntity",
		&returns,
	)
	return returns
}


func NewEmrStudioServiceRole(scope constructs.Construct, name *string, props *EmrStudioServiceRoleProps) EmrStudioServiceRole {
	_init_.Initialize()

	if err := validateNewEmrStudioServiceRoleParameters(scope, name, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrStudioServiceRole{}

	_jsii_.Create(
		"cdk-emrserverless-with-delta-lake.EmrStudioServiceRole",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

func NewEmrStudioServiceRole_Override(e EmrStudioServiceRole, scope constructs.Construct, name *string, props *EmrStudioServiceRoleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-emrserverless-with-delta-lake.EmrStudioServiceRole",
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
func EmrStudioServiceRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmrStudioServiceRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-emrserverless-with-delta-lake.EmrStudioServiceRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudioServiceRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

