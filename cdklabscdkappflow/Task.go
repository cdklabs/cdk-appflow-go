package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A representation of a unitary action on the record fields.
// Experimental.
type Task interface {
	ITask
	// Experimental.
	ConnectorOperator() *TaskConnectorOperator
	// Experimental.
	SetConnectorOperator(val *TaskConnectorOperator)
	// Experimental.
	DestinationField() *string
	// Experimental.
	SetDestinationField(val *string)
	// Experimental.
	Properties() *TaskProperties
	// Experimental.
	SetProperties(val *TaskProperties)
	// Experimental.
	SourceFields() *[]*string
	// Experimental.
	SetSourceFields(val *[]*string)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	Bind(_flow IFlow, source ISource) *awsappflow.CfnFlow_TaskProperty
}

// The jsii proxy struct for Task
type jsiiProxy_Task struct {
	jsiiProxy_ITask
}

func (j *jsiiProxy_Task) ConnectorOperator() *TaskConnectorOperator {
	var returns *TaskConnectorOperator
	_jsii_.Get(
		j,
		"connectorOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) DestinationField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Properties() *TaskProperties {
	var returns *TaskProperties
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) SourceFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewTask(type_ *string, sourceFields *[]*string, connectorOperator *TaskConnectorOperator, properties *TaskProperties, destinationField *string) Task {
	_init_.Initialize()

	if err := validateNewTaskParameters(type_, sourceFields, connectorOperator, properties); err != nil {
		panic(err)
	}
	j := jsiiProxy_Task{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.Task",
		[]interface{}{type_, sourceFields, connectorOperator, properties, destinationField},
		&j,
	)

	return &j
}

// Experimental.
func NewTask_Override(t Task, type_ *string, sourceFields *[]*string, connectorOperator *TaskConnectorOperator, properties *TaskProperties, destinationField *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.Task",
		[]interface{}{type_, sourceFields, connectorOperator, properties, destinationField},
		t,
	)
}

func (j *jsiiProxy_Task)SetConnectorOperator(val *TaskConnectorOperator) {
	if err := j.validateSetConnectorOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorOperator",
		val,
	)
}

func (j *jsiiProxy_Task)SetDestinationField(val *string) {
	_jsii_.Set(
		j,
		"destinationField",
		val,
	)
}

func (j *jsiiProxy_Task)SetProperties(val *TaskProperties) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_Task)SetSourceFields(val *[]*string) {
	if err := j.validateSetSourceFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFields",
		val,
	)
}

func (j *jsiiProxy_Task)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_Task) Bind(_flow IFlow, source ISource) *awsappflow.CfnFlow_TaskProperty {
	if err := t.validateBindParameters(_flow, source); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_TaskProperty

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{_flow, source},
		&returns,
	)

	return returns
}

