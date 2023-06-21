package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// Experimental.
type MarketoSource interface {
	ISource
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy struct for MarketoSource
type jsiiProxy_MarketoSource struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_MarketoSource) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewMarketoSource(props *MarketoSourceProps) MarketoSource {
	_init_.Initialize()

	if err := validateNewMarketoSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_MarketoSource{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.MarketoSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewMarketoSource_Override(m MarketoSource, props *MarketoSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.MarketoSource",
		[]interface{}{props},
		m,
	)
}

func (m *jsiiProxy_MarketoSource) Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
	if err := m.validateBindParameters(flow); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_SourceFlowConfigProperty

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{flow},
		&returns,
	)

	return returns
}

