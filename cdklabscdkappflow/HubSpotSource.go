package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A class that represents a Hubspot Source.
// Experimental.
type HubSpotSource interface {
	ISource
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy struct for HubSpotSource
type jsiiProxy_HubSpotSource struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_HubSpotSource) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewHubSpotSource(props *HubSpotSourceProps) HubSpotSource {
	_init_.Initialize()

	if err := validateNewHubSpotSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HubSpotSource{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.HubSpotSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewHubSpotSource_Override(h HubSpotSource, props *HubSpotSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.HubSpotSource",
		[]interface{}{props},
		h,
	)
}

func (h *jsiiProxy_HubSpotSource) Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
	if err := h.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_SourceFlowConfigProperty

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

