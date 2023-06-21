package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// Experimental.
type S3Destination interface {
	IDestination
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(flow IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty
}

// The jsii proxy struct for S3Destination
type jsiiProxy_S3Destination struct {
	jsiiProxy_IDestination
}

func (j *jsiiProxy_S3Destination) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewS3Destination(props *S3DestinationProps) S3Destination {
	_init_.Initialize()

	if err := validateNewS3DestinationParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3Destination{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.S3Destination",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewS3Destination_Override(s S3Destination, props *S3DestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.S3Destination",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_S3Destination) Bind(flow IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty {
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

