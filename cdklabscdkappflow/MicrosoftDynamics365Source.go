package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A class that represents a Microsoft Dynamics 365 Source.
// Experimental.
type MicrosoftDynamics365Source interface {
	ISource
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy struct for MicrosoftDynamics365Source
type jsiiProxy_MicrosoftDynamics365Source struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_MicrosoftDynamics365Source) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewMicrosoftDynamics365Source(props *MicrosoftDynamics365SourceProps) MicrosoftDynamics365Source {
	_init_.Initialize()

	if err := validateNewMicrosoftDynamics365SourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_MicrosoftDynamics365Source{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365Source",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewMicrosoftDynamics365Source_Override(m MicrosoftDynamics365Source, props *MicrosoftDynamics365SourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365Source",
		[]interface{}{props},
		m,
	)
}

func (m *jsiiProxy_MicrosoftDynamics365Source) Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
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

