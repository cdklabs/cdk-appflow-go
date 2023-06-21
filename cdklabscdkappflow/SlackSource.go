package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// Experimental.
type SlackSource interface {
	ISource
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy struct for SlackSource
type jsiiProxy_SlackSource struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_SlackSource) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewSlackSource(props *SlackSourceProps) SlackSource {
	_init_.Initialize()

	if err := validateNewSlackSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SlackSource{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.SlackSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewSlackSource_Override(s SlackSource, props *SlackSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.SlackSource",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_SlackSource) Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
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

