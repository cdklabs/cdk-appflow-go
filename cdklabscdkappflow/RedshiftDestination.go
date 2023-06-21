package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// Experimental.
type RedshiftDestination interface {
	IDestination
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(scope IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty
}

// The jsii proxy struct for RedshiftDestination
type jsiiProxy_RedshiftDestination struct {
	jsiiProxy_IDestination
}

func (j *jsiiProxy_RedshiftDestination) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewRedshiftDestination(props *RedshiftDestinationProps) RedshiftDestination {
	_init_.Initialize()

	if err := validateNewRedshiftDestinationParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedshiftDestination{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.RedshiftDestination",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewRedshiftDestination_Override(r RedshiftDestination, props *RedshiftDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.RedshiftDestination",
		[]interface{}{props},
		r,
	)
}

func (r *jsiiProxy_RedshiftDestination) Bind(scope IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty {
	if err := r.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_DestinationFlowConfigProperty

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

