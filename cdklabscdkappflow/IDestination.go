package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A destination of an AppFlow flow.
// Experimental.
type IDestination interface {
	IVertex
	// Experimental.
	Bind(scope IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty
}

// The jsii proxy for IDestination
type jsiiProxy_IDestination struct {
	jsiiProxy_IVertex
}

func (i *jsiiProxy_IDestination) Bind(scope IFlow) *awsappflow.CfnFlow_DestinationFlowConfigProperty {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_DestinationFlowConfigProperty

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

