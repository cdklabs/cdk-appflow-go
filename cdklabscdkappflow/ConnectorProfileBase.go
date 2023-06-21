package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/internal"
)

// Experimental.
type ConnectorProfileBase interface {
	awscdk.Resource
	IConnectorProfile
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
	BuildConnectorProfileProperties(props *ConnectorProfileProps) *awsappflow.CfnConnectorProfile_ConnectorProfilePropertiesProperty
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

// The jsii proxy struct for ConnectorProfileBase
type jsiiProxy_ConnectorProfileBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IConnectorProfile
}

func (j *jsiiProxy_ConnectorProfileBase) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectorProfileBase) Credentials() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectorProfileBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectorProfileBase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectorProfileBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectorProfileBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectorProfileBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewConnectorProfileBase_Override(c ConnectorProfileBase, scope constructs.Construct, id *string, props *ConnectorProfileProps, connectorType ConnectorType) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.ConnectorProfileBase",
		[]interface{}{scope, id, props, connectorType},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ConnectorProfileBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConnectorProfileBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.ConnectorProfileBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func ConnectorProfileBase_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateConnectorProfileBase_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.ConnectorProfileBase",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ConnectorProfileBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateConnectorProfileBase_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.ConnectorProfileBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorProfileBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := c.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_ConnectorProfileBase) BuildConnectorProfileCredentials(props *ConnectorProfileProps) *awsappflow.CfnConnectorProfile_ConnectorProfileCredentialsProperty {
	if err := c.validateBuildConnectorProfileCredentialsParameters(props); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnConnectorProfile_ConnectorProfileCredentialsProperty

	_jsii_.Invoke(
		c,
		"buildConnectorProfileCredentials",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorProfileBase) BuildConnectorProfileProperties(props *ConnectorProfileProps) *awsappflow.CfnConnectorProfile_ConnectorProfilePropertiesProperty {
	if err := c.validateBuildConnectorProfilePropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnConnectorProfile_ConnectorProfilePropertiesProperty

	_jsii_.Invoke(
		c,
		"buildConnectorProfileProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorProfileBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorProfileBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := c.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorProfileBase) GetResourceNameAttribute(nameAttr *string) *string {
	if err := c.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorProfileBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorProfileBase) TryAddNodeDependency(scope constructs.IConstruct, resource interface{}) {
	if err := c.validateTryAddNodeDependencyParameters(scope, resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"tryAddNodeDependency",
		[]interface{}{scope, resource},
	)
}

