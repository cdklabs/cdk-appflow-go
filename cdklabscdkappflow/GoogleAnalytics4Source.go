package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A class that represents a Google Analytics v4 Source.
// Experimental.
type GoogleAnalytics4Source interface {
	ISource
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy struct for GoogleAnalytics4Source
type jsiiProxy_GoogleAnalytics4Source struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_GoogleAnalytics4Source) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewGoogleAnalytics4Source(props *GoogleAnalytics4SourceProps) GoogleAnalytics4Source {
	_init_.Initialize()

	if err := validateNewGoogleAnalytics4SourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAnalytics4Source{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.GoogleAnalytics4Source",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewGoogleAnalytics4Source_Override(g GoogleAnalytics4Source, props *GoogleAnalytics4SourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.GoogleAnalytics4Source",
		[]interface{}{props},
		g,
	)
}

func (g *jsiiProxy_GoogleAnalytics4Source) Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
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

