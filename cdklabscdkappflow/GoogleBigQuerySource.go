package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A class that represents a Google BigQuery Source.
// Experimental.
type GoogleBigQuerySource interface {
	ISource
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy struct for GoogleBigQuerySource
type jsiiProxy_GoogleBigQuerySource struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_GoogleBigQuerySource) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewGoogleBigQuerySource(props *GoogleBigQuerySourceProps) GoogleBigQuerySource {
	_init_.Initialize()

	if err := validateNewGoogleBigQuerySourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigQuerySource{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.GoogleBigQuerySource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewGoogleBigQuerySource_Override(g GoogleBigQuerySource, props *GoogleBigQuerySourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.GoogleBigQuerySource",
		[]interface{}{props},
		g,
	)
}

func (g *jsiiProxy_GoogleBigQuerySource) Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
	if err := g.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_SourceFlowConfigProperty

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

