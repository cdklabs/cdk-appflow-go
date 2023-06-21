package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// Experimental.
type SalesforceDestination interface {
	IDestination
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(flow IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty
}

// The jsii proxy struct for SalesforceDestination
type jsiiProxy_SalesforceDestination struct {
	jsiiProxy_IDestination
}

func (j *jsiiProxy_SalesforceDestination) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewSalesforceDestination(props *SalesforceDestinationProps) SalesforceDestination {
	_init_.Initialize()

	if err := validateNewSalesforceDestinationParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SalesforceDestination{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.SalesforceDestination",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewSalesforceDestination_Override(s SalesforceDestination, props *SalesforceDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.SalesforceDestination",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_SalesforceDestination) Bind(flow IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty {
	if err := s.validateBindParameters(flow); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_DestinationFlowConfigProperty

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{flow},
		&returns,
	)

	return returns
}

