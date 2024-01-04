package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// A base class for triggered flows.
// Experimental.
type TriggeredFlowBase interface {
	FlowBase
	IFlow
	// The ARN of the flow.
	// Experimental.
	Arn() *string
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
	// The name of the flow.
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
	// The type of the flow.
	// Experimental.
	Type() FlowType
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
	// Experimental.
	Metric(metricName *string, options *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Creates a metric to report the number of records that Amazon AppFlow attempted to transfer for the flow run.
	// Experimental.
	MetricFlowExecutionRecordsProcessed(options *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Creates a metric to report the number of failed flow runs.
	// Experimental.
	MetricFlowExecutionsFailed(options *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Creates a metric to report the number of flow runs started.
	// Experimental.
	MetricFlowExecutionsStarted(options *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Creates a metric to report the number of successful flow runs.
	// Experimental.
	MetricFlowExecutionsSucceeded(options *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Creates a metric to report the  interval, in milliseconds, between the time the flow starts and the time it finishes.
	// Experimental.
	MetricFlowExecutionTime(options *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	OnDeactivated(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Experimental.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Experimental.
	OnRunCompleted(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Experimental.
	OnRunStarted(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TriggeredFlowBase
type jsiiProxy_TriggeredFlowBase struct {
	jsiiProxy_FlowBase
	jsiiProxy_IFlow
}

func (j *jsiiProxy_TriggeredFlowBase) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TriggeredFlowBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TriggeredFlowBase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TriggeredFlowBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TriggeredFlowBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TriggeredFlowBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TriggeredFlowBase) Type() FlowType {
	var returns FlowType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewTriggeredFlowBase_Override(t TriggeredFlowBase, scope constructs.Construct, id *string, props *FlowBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.TriggeredFlowBase",
		[]interface{}{scope, id, props},
		t,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func TriggeredFlowBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTriggeredFlowBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.TriggeredFlowBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func TriggeredFlowBase_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateTriggeredFlowBase_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.TriggeredFlowBase",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func TriggeredFlowBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateTriggeredFlowBase_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.TriggeredFlowBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Experimental.
func TriggeredFlowBase_SetStatus(autoActivate *bool, status FlowStatus) FlowStatus {
	_init_.Initialize()

	var returns FlowStatus

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.TriggeredFlowBase",
		"setStatus",
		[]interface{}{autoActivate, status},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TriggeredFlowBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := t.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (t *jsiiProxy_TriggeredFlowBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TriggeredFlowBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := t.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TriggeredFlowBase) GetResourceNameAttribute(nameAttr *string) *string {
	if err := t.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TriggeredFlowBase) Metric(metricName *string, options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricParameters(metricName, options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metric",
		[]interface{}{metricName, options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TriggeredFlowBase) MetricFlowExecutionRecordsProcessed(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricFlowExecutionRecordsProcessedParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricFlowExecutionRecordsProcessed",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TriggeredFlowBase) MetricFlowExecutionsFailed(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricFlowExecutionsFailedParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricFlowExecutionsFailed",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TriggeredFlowBase) MetricFlowExecutionsStarted(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricFlowExecutionsStartedParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricFlowExecutionsStarted",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TriggeredFlowBase) MetricFlowExecutionsSucceeded(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricFlowExecutionsSucceededParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricFlowExecutionsSucceeded",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TriggeredFlowBase) MetricFlowExecutionTime(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricFlowExecutionTimeParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricFlowExecutionTime",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TriggeredFlowBase) OnDeactivated(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := t.validateOnDeactivatedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		t,
		"onDeactivated",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TriggeredFlowBase) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := t.validateOnEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		t,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TriggeredFlowBase) OnRunCompleted(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := t.validateOnRunCompletedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		t,
		"onRunCompleted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TriggeredFlowBase) OnRunStarted(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := t.validateOnRunStartedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		t,
		"onRunStarted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TriggeredFlowBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

