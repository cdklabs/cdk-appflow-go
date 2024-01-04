package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// Experimental.
type JdbcSmallDataScaleSource interface {
	ISource
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy struct for JdbcSmallDataScaleSource
type jsiiProxy_JdbcSmallDataScaleSource struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_JdbcSmallDataScaleSource) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewJdbcSmallDataScaleSource(props *JdbcSmallDataScaleSourceProps) JdbcSmallDataScaleSource {
	_init_.Initialize()

	if err := validateNewJdbcSmallDataScaleSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_JdbcSmallDataScaleSource{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.JdbcSmallDataScaleSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewJdbcSmallDataScaleSource_Override(j JdbcSmallDataScaleSource, props *JdbcSmallDataScaleSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.JdbcSmallDataScaleSource",
		[]interface{}{props},
		j,
	)
}

func (j *jsiiProxy_JdbcSmallDataScaleSource) Bind(flow IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
	if err := j.validateBindParameters(flow); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_SourceFlowConfigProperty

	_jsii_.Invoke(
		j,
		"bind",
		[]interface{}{flow},
		&returns,
	)

	return returns
}

