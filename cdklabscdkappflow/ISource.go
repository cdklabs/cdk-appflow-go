package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A source of an AppFlow flow.
// Experimental.
type ISource interface {
	IVertex
	// Experimental.
	Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy for ISource
type jsiiProxy_ISource struct {
	jsiiProxy_IVertex
}

func (i *jsiiProxy_ISource) Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_SourceFlowConfigProperty

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

