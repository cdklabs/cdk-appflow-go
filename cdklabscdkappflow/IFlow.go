package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/internal"
)

// Experimental.
type IFlow interface {
	awscdk.IResource
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

func (i *jsiiProxy_IFlow) MetricFlowExecutionRecordsProcessed(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricFlowExecutionRecordsProcessedParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFlowExecutionRecordsProcessed",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFlow) MetricFlowExecutionsFailed(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricFlowExecutionsFailedParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFlowExecutionsFailed",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFlow) MetricFlowExecutionsStarted(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricFlowExecutionsStartedParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFlowExecutionsStarted",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFlow) MetricFlowExecutionsSucceeded(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricFlowExecutionsSucceededParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFlowExecutionsSucceeded",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFlow) MetricFlowExecutionTime(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricFlowExecutionTimeParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFlowExecutionTime",
		[]interface{}{options},
		&returns,
	)

	return returns
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

