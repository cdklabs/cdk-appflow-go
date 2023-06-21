package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// Experimental.
type EventBridgeDestination interface {
	IDestination
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(flow IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty
}

// The jsii proxy struct for EventBridgeDestination
type jsiiProxy_EventBridgeDestination struct {
	jsiiProxy_IDestination
}

func (j *jsiiProxy_EventBridgeDestination) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewEventBridgeDestination(props *EventBridgeDestinationProps) EventBridgeDestination {
	_init_.Initialize()

	if err := validateNewEventBridgeDestinationParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventBridgeDestination{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.EventBridgeDestination",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewEventBridgeDestination_Override(e EventBridgeDestination, props *EventBridgeDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.EventBridgeDestination",
		[]interface{}{props},
		e,
	)
}

func (e *jsiiProxy_EventBridgeDestination) Bind(flow IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty {
	if err := e.validateBindParameters(flow); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_DestinationFlowConfigProperty

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{flow},
		&returns,
	)

	return returns
}

