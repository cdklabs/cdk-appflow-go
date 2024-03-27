package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A class that represents a Mailchimp v3 Source.
// Experimental.
type MailchimpSource interface {
	ISource
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy struct for MailchimpSource
type jsiiProxy_MailchimpSource struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_MailchimpSource) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewMailchimpSource(props *MailchimpSourceProps) MailchimpSource {
	_init_.Initialize()

	if err := validateNewMailchimpSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_MailchimpSource{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.MailchimpSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewMailchimpSource_Override(m MailchimpSource, props *MailchimpSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.MailchimpSource",
		[]interface{}{props},
		m,
	)
}

func (m *jsiiProxy_MailchimpSource) Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
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

