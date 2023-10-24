package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A Snowflake destination.
// Experimental.
type SnowflakeDestination interface {
	IDestination
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(scope IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty
}

// The jsii proxy struct for SnowflakeDestination
type jsiiProxy_SnowflakeDestination struct {
	jsiiProxy_IDestination
}

func (j *jsiiProxy_SnowflakeDestination) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewSnowflakeDestination(props *SnowflakeDestinationProps) SnowflakeDestination {
	_init_.Initialize()

	if err := validateNewSnowflakeDestinationParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnowflakeDestination{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.SnowflakeDestination",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewSnowflakeDestination_Override(s SnowflakeDestination, props *SnowflakeDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.SnowflakeDestination",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_SnowflakeDestination) Bind(scope IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty {
	if err := s.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_DestinationFlowConfigProperty

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

