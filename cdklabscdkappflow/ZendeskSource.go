package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// Experimental.
type ZendeskSource interface {
	ISource
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy struct for ZendeskSource
type jsiiProxy_ZendeskSource struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_ZendeskSource) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewZendeskSource(props *ZendeskSourceProps) ZendeskSource {
	_init_.Initialize()

	if err := validateNewZendeskSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZendeskSource{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.ZendeskSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewZendeskSource_Override(z ZendeskSource, props *ZendeskSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.ZendeskSource",
		[]interface{}{props},
		z,
	)
}

func (z *jsiiProxy_ZendeskSource) Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
	if err := z.validateBindParameters(flow); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_SourceFlowConfigProperty

	_jsii_.Invoke(
		z,
		"bind",
		[]interface{}{flow},
		&returns,
	)

	return returns
}

