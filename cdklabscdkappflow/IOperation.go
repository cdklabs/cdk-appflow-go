package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A representation of a set of tasks that deliver complete operation.
// Experimental.
type IOperation interface {
	// Experimental.
	Bind(flow IFlow, source ISource) *[]*awsappflow.CfnFlow_TaskProperty
}

// The jsii proxy for IOperation
type jsiiProxy_IOperation struct {
	_ byte // padding
}

func (i *jsiiProxy_IOperation) Bind(flow IFlow, source ISource) *[]*awsappflow.CfnFlow_TaskProperty {
	if err := i.validateBindParameters(flow, source); err != nil {
		panic(err)
	}
	var returns *[]*awsappflow.CfnFlow_TaskProperty

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{flow, source},
		&returns,
	)

	return returns
}

