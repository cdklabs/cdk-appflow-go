package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// Experimental.
type HubSpotDestination interface {
	IDestination
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(flow IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty
}

// The jsii proxy struct for HubSpotDestination
type jsiiProxy_HubSpotDestination struct {
	jsiiProxy_IDestination
}

func (j *jsiiProxy_HubSpotDestination) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewHubSpotDestination(props *HubSpotDestinationProps) HubSpotDestination {
	_init_.Initialize()

	if err := validateNewHubSpotDestinationParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HubSpotDestination{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.HubSpotDestination",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewHubSpotDestination_Override(h HubSpotDestination, props *HubSpotDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.HubSpotDestination",
		[]interface{}{props},
		h,
	)
}

func (h *jsiiProxy_HubSpotDestination) Bind(flow IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty {
	if err := h.validateBindParameters(flow); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_DestinationFlowConfigProperty

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{flow},
		&returns,
	)

	return returns
}

