package cdkemrserverlesswithdeltalake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-emrserverless-with-delta-lake-go/cdkemrserverlesswithdeltalake/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-emrserverless-with-delta-lake-go/cdkemrserverlesswithdeltalake/v2/internal"
)

// Created an engine security group for EMR notebooks.
//
// For detail, plrease refer to [Engine security group](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-studio-security-groups.html#emr-studio-security-group-instructions).
//
// ```ts
// const workSpaceSecurityGroup = new EmrStudioWorkspaceSecurityGroup(this, 'Workspace', { vpc: baseVpc });
// const engineSecurityGroup = new EmrStudioEngineSecurityGroup(this, 'Engine', { vpc: baseVpc });
// workSpaceSecurityGroup.entity.connections.allowTo(engineSecurityGroup.entity, ec2.Port.tcp(18888), 'Allow traffic to any resources in the Engine security group for EMR Studio.');
// workSpaceSecurityGroup.entity.addEgressRule(ec2.Peer.anyIpv4(), ec2.Port.tcp(443), 'Allow traffic to the internet to link publicly hosted Git repositories to Workspaces.');
// ```.
type EmrStudioEngineSecurityGroup interface {
	constructs.Construct
	// The representative of the security group as the EMR Studio engine security group.
	Entity() awsec2.SecurityGroup
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for EmrStudioEngineSecurityGroup
type jsiiProxy_EmrStudioEngineSecurityGroup struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EmrStudioEngineSecurityGroup) Entity() awsec2.SecurityGroup {
	var returns awsec2.SecurityGroup
	_jsii_.Get(
		j,
		"entity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioEngineSecurityGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewEmrStudioEngineSecurityGroup(scope constructs.Construct, name *string, props *EmrStudioEngineSecurityGroupProps) EmrStudioEngineSecurityGroup {
	_init_.Initialize()

	if err := validateNewEmrStudioEngineSecurityGroupParameters(scope, name, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrStudioEngineSecurityGroup{}

	_jsii_.Create(
		"cdk-emrserverless-with-delta-lake.EmrStudioEngineSecurityGroup",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

func NewEmrStudioEngineSecurityGroup_Override(e EmrStudioEngineSecurityGroup, scope constructs.Construct, name *string, props *EmrStudioEngineSecurityGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-emrserverless-with-delta-lake.EmrStudioEngineSecurityGroup",
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
func EmrStudioEngineSecurityGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmrStudioEngineSecurityGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-emrserverless-with-delta-lake.EmrStudioEngineSecurityGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudioEngineSecurityGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

