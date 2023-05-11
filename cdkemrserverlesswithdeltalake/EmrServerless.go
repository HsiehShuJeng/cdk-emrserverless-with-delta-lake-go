package cdkemrserverlesswithdeltalake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-emrserverless-with-delta-lake-go/cdkemrserverlesswithdeltalake/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-emrserverless-with-delta-lake-go/cdkemrserverlesswithdeltalake/v2/internal"
)

// Creates an EMR Studio, an EMR cluster template for the studio, and an EMR Serverless application.
//
// ```ts
// // the quickiest deployment
// new EmrServerless(this, 'EmrServerless');
//
// // custom deployment references
// new EmrServerless(this, 'EmrServerless', {
//     vpcId: 'vpc-idididid',
// });
//
// new EmrServerless(this, 'EmrServerless', {
//     vpcId: 'vpc-idididid',
//     subnetIds: ['subnet-eeeee', 'subnet-fffff']
// });
//
// const myRole = new iam.Role.fromRoleName('MyRole');
// new EmrServerless(this, 'EmrServerless', {
//     serviceCatalogProps: {
//         role: myRole
//     }
// });
//
// const myUser = new iam.Role.fromUserName('MyUser');
// new EmrServerless(this, 'EmrServerless', {
//     vpcId: 'vpc-idididid',
//     subnetIds: ['subnet-eeeee', 'subnet-fffff'],
//     serviceCatalogProps: {
//         user: myUser
//     }
// });
//
// const myGroup = new iam.Group.fromGroupName('MyGroup');
// new EmrServerless(this, 'EmrServerless', {
//     serviceCatalogProps: {
//         group: myGroup
//     }
// });
// ```.
type EmrServerless interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for EmrServerless
type jsiiProxy_EmrServerless struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EmrServerless) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewEmrServerless(scope constructs.Construct, name *string, props *EmrServerlessProps) EmrServerless {
	_init_.Initialize()

	if err := validateNewEmrServerlessParameters(scope, name, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrServerless{}

	_jsii_.Create(
		"cdk-emrserverless-with-delta-lake.EmrServerless",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

func NewEmrServerless_Override(e EmrServerless, scope constructs.Construct, name *string, props *EmrServerlessProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-emrserverless-with-delta-lake.EmrServerless",
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
func EmrServerless_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmrServerless_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-emrserverless-with-delta-lake.EmrServerless",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrServerless) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

