package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A class that represents a Google Ads v4 Source.
// Experimental.
type GoogleAdsSource interface {
	ISource
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy struct for GoogleAdsSource
type jsiiProxy_GoogleAdsSource struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_GoogleAdsSource) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewGoogleAdsSource(props *GoogleAdsSourceProps) GoogleAdsSource {
	_init_.Initialize()

	if err := validateNewGoogleAdsSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAdsSource{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.GoogleAdsSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewGoogleAdsSource_Override(g GoogleAdsSource, props *GoogleAdsSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.GoogleAdsSource",
		[]interface{}{props},
		g,
	)
}

func (g *jsiiProxy_GoogleAdsSource) Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
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

