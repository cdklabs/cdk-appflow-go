package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// Experimental.
type SAPOdataSource interface {
	ISource
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy struct for SAPOdataSource
type jsiiProxy_SAPOdataSource struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_SAPOdataSource) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewSAPOdataSource(props *SAPOdataSourceProps) SAPOdataSource {
	_init_.Initialize()

	if err := validateNewSAPOdataSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SAPOdataSource{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.SAPOdataSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewSAPOdataSource_Override(s SAPOdataSource, props *SAPOdataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.SAPOdataSource",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_SAPOdataSource) Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
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

