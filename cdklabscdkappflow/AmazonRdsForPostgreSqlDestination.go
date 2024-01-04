package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// Represents a destination for the Amazon RDS for PostgreSQL connector.
// Experimental.
type AmazonRdsForPostgreSqlDestination interface {
	IDestination
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(flow IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty
}

// The jsii proxy struct for AmazonRdsForPostgreSqlDestination
type jsiiProxy_AmazonRdsForPostgreSqlDestination struct {
	jsiiProxy_IDestination
}

func (j *jsiiProxy_AmazonRdsForPostgreSqlDestination) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Creates a new instance of the AmazonRdsForPostgreSqlDestination.
// Experimental.
func NewAmazonRdsForPostgreSqlDestination(props *AmazonRdsForPostgreSqlDestinationProps) AmazonRdsForPostgreSqlDestination {
	_init_.Initialize()

	if err := validateNewAmazonRdsForPostgreSqlDestinationParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AmazonRdsForPostgreSqlDestination{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.AmazonRdsForPostgreSqlDestination",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Creates a new instance of the AmazonRdsForPostgreSqlDestination.
// Experimental.
func NewAmazonRdsForPostgreSqlDestination_Override(a AmazonRdsForPostgreSqlDestination, props *AmazonRdsForPostgreSqlDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.AmazonRdsForPostgreSqlDestination",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AmazonRdsForPostgreSqlDestination) Bind(flow IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty {
	if err := a.validateBindParameters(flow); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_DestinationFlowConfigProperty

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{flow},
		&returns,
	)

	return returns
}

