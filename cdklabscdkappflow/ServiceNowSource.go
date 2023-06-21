package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// Experimental.
type ServiceNowSource interface {
	ISource
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy struct for ServiceNowSource
type jsiiProxy_ServiceNowSource struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_ServiceNowSource) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewServiceNowSource(props *ServiceNowSourceProps) ServiceNowSource {
	_init_.Initialize()

	if err := validateNewServiceNowSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceNowSource{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.ServiceNowSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewServiceNowSource_Override(s ServiceNowSource, props *ServiceNowSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.ServiceNowSource",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_ServiceNowSource) Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
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

