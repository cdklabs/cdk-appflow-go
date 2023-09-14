package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/internal"
)

// Experimental.
type IFlow interface {
	awscdk.IResource
	// Experimental.
	OnRunCompleted(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Experimental.
	OnRunStarted(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// The ARN of the flow.
	// Experimental.
	Arn() *string
	// The name of the flow.
	// Experimental.
	Name() *string
	// The type of the flow.
	// Experimental.
	Type() FlowType
}

// The jsii proxy for IFlow
type jsiiProxy_IFlow struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IFlow) OnRunCompleted(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnRunCompletedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onRunCompleted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFlow) OnRunStarted(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnRunStartedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onRunStarted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IFlow) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlow) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlow) Type() FlowType {
	var returns FlowType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

