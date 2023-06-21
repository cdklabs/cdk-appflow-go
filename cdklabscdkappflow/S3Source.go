package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// Experimental.
type S3Source interface {
	ISource
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy struct for S3Source
type jsiiProxy_S3Source struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_S3Source) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewS3Source(props *S3SourceProps) S3Source {
	_init_.Initialize()

	if err := validateNewS3SourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3Source{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.S3Source",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewS3Source_Override(s S3Source, props *S3SourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.S3Source",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_S3Source) Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
	if err := s.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_SourceFlowConfigProperty

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

