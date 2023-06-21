package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A class that represents a Google Analytics v4 Source.
// Experimental.
type MicrosoftSharepointOnlineSource interface {
	ISource
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy struct for MicrosoftSharepointOnlineSource
type jsiiProxy_MicrosoftSharepointOnlineSource struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_MicrosoftSharepointOnlineSource) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewMicrosoftSharepointOnlineSource(props *MicrosoftSharepointOnlineSourceProps) MicrosoftSharepointOnlineSource {
	_init_.Initialize()

	if err := validateNewMicrosoftSharepointOnlineSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_MicrosoftSharepointOnlineSource{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.MicrosoftSharepointOnlineSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewMicrosoftSharepointOnlineSource_Override(m MicrosoftSharepointOnlineSource, props *MicrosoftSharepointOnlineSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.MicrosoftSharepointOnlineSource",
		[]interface{}{props},
		m,
	)
}

func (m *jsiiProxy_MicrosoftSharepointOnlineSource) Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
	if err := m.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_SourceFlowConfigProperty

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

