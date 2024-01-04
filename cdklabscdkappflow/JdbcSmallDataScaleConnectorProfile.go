package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
)

// The connector profile for the JDBC connector.
// Experimental.
type JdbcSmallDataScaleConnectorProfile interface {
	ConnectorProfileBase
	// Experimental.
	Arn() *string
	// Experimental.
	Credentials() awssecretsmanager.ISecret
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	BuildConnectorProfileCredentials(props *ConnectorProfileProps) *awsappflow.CfnConnectorProfile_ConnectorProfileCredentialsProperty
	// Experimental.
	BuildConnectorProfileProperties(_props *ConnectorProfileProps) *awsappflow.CfnConnectorProfile_ConnectorProfilePropertiesProperty
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Experimental.
	TryAddNodeDependency(scope constructs.IConstruct, resource interface{})
}

// The jsii proxy struct for JdbcSmallDataScaleConnectorProfile
type jsiiProxy_JdbcSmallDataScaleConnectorProfile struct {
	jsiiProxy_ConnectorProfileBase
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) Credentials() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Creates a new instance of the JdbcSmallDataScaleConnectorProfile.
// Experimental.
func NewJdbcSmallDataScaleConnectorProfile(scope constructs.Construct, id *string, props *JdbcSmallDataScaleConnectorProfileProps) JdbcSmallDataScaleConnectorProfile {
	_init_.Initialize()

	if err := validateNewJdbcSmallDataScaleConnectorProfileParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_JdbcSmallDataScaleConnectorProfile{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.JdbcSmallDataScaleConnectorProfile",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new instance of the JdbcSmallDataScaleConnectorProfile.
// Experimental.
func NewJdbcSmallDataScaleConnectorProfile_Override(j JdbcSmallDataScaleConnectorProfile, scope constructs.Construct, id *string, props *JdbcSmallDataScaleConnectorProfileProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.JdbcSmallDataScaleConnectorProfile",
		[]interface{}{scope, id, props},
		j,
	)
}

// Imports an existing JdbcSmallDataScaleConnectorProfile.
//
// Returns: An instance of the JdbcSmallDataScaleConnectorProfile.
// Experimental.
func JdbcSmallDataScaleConnectorProfile_FromConnectionProfileArn(scope constructs.Construct, id *string, arn *string) JdbcSmallDataScaleConnectorProfile {
	_init_.Initialize()

	if err := validateJdbcSmallDataScaleConnectorProfile_FromConnectionProfileArnParameters(scope, id, arn); err != nil {
		panic(err)
	}
	var returns JdbcSmallDataScaleConnectorProfile

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.JdbcSmallDataScaleConnectorProfile",
		"fromConnectionProfileArn",
		[]interface{}{scope, id, arn},
		&returns,
	)

	return returns
}

// Imports an existing JdbcSmallDataScaleConnectorProfile.
//
// Returns: An instance of the JdbcSmallDataScaleConnectorProfile.
// Experimental.
func JdbcSmallDataScaleConnectorProfile_FromConnectionProfileName(scope constructs.Construct, id *string, name *string) JdbcSmallDataScaleConnectorProfile {
	_init_.Initialize()

	if err := validateJdbcSmallDataScaleConnectorProfile_FromConnectionProfileNameParameters(scope, id, name); err != nil {
		panic(err)
	}
	var returns JdbcSmallDataScaleConnectorProfile

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.JdbcSmallDataScaleConnectorProfile",
		"fromConnectionProfileName",
		[]interface{}{scope, id, name},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func JdbcSmallDataScaleConnectorProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJdbcSmallDataScaleConnectorProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.JdbcSmallDataScaleConnectorProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func JdbcSmallDataScaleConnectorProfile_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateJdbcSmallDataScaleConnectorProfile_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.JdbcSmallDataScaleConnectorProfile",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func JdbcSmallDataScaleConnectorProfile_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateJdbcSmallDataScaleConnectorProfile_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.JdbcSmallDataScaleConnectorProfile",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := j.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) BuildConnectorProfileCredentials(props *ConnectorProfileProps) *awsappflow.CfnConnectorProfile_ConnectorProfileCredentialsProperty {
	if err := j.validateBuildConnectorProfileCredentialsParameters(props); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnConnectorProfile_ConnectorProfileCredentialsProperty

	_jsii_.Invoke(
		j,
		"buildConnectorProfileCredentials",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) BuildConnectorProfileProperties(_props *ConnectorProfileProps) *awsappflow.CfnConnectorProfile_ConnectorProfilePropertiesProperty {
	if err := j.validateBuildConnectorProfilePropertiesParameters(_props); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnConnectorProfile_ConnectorProfilePropertiesProperty

	_jsii_.Invoke(
		j,
		"buildConnectorProfileProperties",
		[]interface{}{_props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := j.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) GetResourceNameAttribute(nameAttr *string) *string {
	if err := j.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) TryAddNodeDependency(scope constructs.IConstruct, resource interface{}) {
	if err := j.validateTryAddNodeDependencyParameters(scope, resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"tryAddNodeDependency",
		[]interface{}{scope, resource},
	)
}

