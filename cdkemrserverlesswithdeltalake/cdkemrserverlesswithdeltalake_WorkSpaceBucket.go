// A construct for the quick demo of EMR Serverless.
package cdkemrserverlesswithdeltalake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-emrserverless-with-delta-lake-go/cdkemrserverlesswithdeltalake/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-emrserverless-with-delta-lake-go/cdkemrserverlesswithdeltalake/v2/internal"
)

type WorkSpaceBucket interface {
	constructs.Construct
	BucketEntity() awss3.Bucket
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for WorkSpaceBucket
type jsiiProxy_WorkSpaceBucket struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_WorkSpaceBucket) BucketEntity() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"bucketEntity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkSpaceBucket) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewWorkSpaceBucket(scope constructs.Construct, name *string, props *WorkSpaceBucketProps) WorkSpaceBucket {
	_init_.Initialize()

	j := jsiiProxy_WorkSpaceBucket{}

	_jsii_.Create(
		"cdk-emrserverless-with-delta-lake.WorkSpaceBucket",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

func NewWorkSpaceBucket_Override(w WorkSpaceBucket, scope constructs.Construct, name *string, props *WorkSpaceBucketProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-emrserverless-with-delta-lake.WorkSpaceBucket",
		[]interface{}{scope, name, props},
		w,
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
func WorkSpaceBucket_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-emrserverless-with-delta-lake.WorkSpaceBucket",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkSpaceBucket) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

