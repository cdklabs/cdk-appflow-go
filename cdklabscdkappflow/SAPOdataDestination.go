package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// Experimental.
type SAPOdataDestination interface {
	IDestination
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(flow IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty
}

// The jsii proxy struct for SAPOdataDestination
type jsiiProxy_SAPOdataDestination struct {
	jsiiProxy_IDestination
}

func (j *jsiiProxy_SAPOdataDestination) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewSAPOdataDestination(props *SAPOdataDestinationProps) SAPOdataDestination {
	_init_.Initialize()

	if err := validateNewSAPOdataDestinationParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SAPOdataDestination{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.SAPOdataDestination",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewSAPOdataDestination_Override(s SAPOdataDestination, props *SAPOdataDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.SAPOdataDestination",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_SAPOdataDestination) Bind(flow IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty {
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

