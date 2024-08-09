package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A class that represents a Asana v3 Source.
// Experimental.
type AsanaSource interface {
	ISource
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy struct for AsanaSource
type jsiiProxy_AsanaSource struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_AsanaSource) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewAsanaSource(props *AsanaSourceProps) AsanaSource {
	_init_.Initialize()

	if err := validateNewAsanaSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AsanaSource{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.AsanaSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAsanaSource_Override(a AsanaSource, props *AsanaSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.AsanaSource",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AsanaSource) Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
	if err := a.validateBindParameters(flow); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_SourceFlowConfigProperty

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{flow},
		&returns,
	)

	return returns
}

