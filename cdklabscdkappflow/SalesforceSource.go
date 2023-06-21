package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// Experimental.
type SalesforceSource interface {
	ISource
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy struct for SalesforceSource
type jsiiProxy_SalesforceSource struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_SalesforceSource) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewSalesforceSource(props *SalesforceSourceProps) SalesforceSource {
	_init_.Initialize()

	if err := validateNewSalesforceSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SalesforceSource{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.SalesforceSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewSalesforceSource_Override(s SalesforceSource, props *SalesforceSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.SalesforceSource",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_SalesforceSource) Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
	if err := s.validateBindParameters(flow); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_SourceFlowConfigProperty

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{flow},
		&returns,
	)

	return returns
}

