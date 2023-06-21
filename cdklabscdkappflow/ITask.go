package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A representation of a unitary action on the record fields.
// Experimental.
type ITask interface {
	// Experimental.
	Bind(flow IFlow, source ISource) *awsappflow.CfnFlow_TaskProperty
}

// The jsii proxy for ITask
type jsiiProxy_ITask struct {
	_ byte // padding
}

func (i *jsiiProxy_ITask) Bind(flow IFlow, source ISource) *awsappflow.CfnFlow_TaskProperty {
	if err := i.validateBindParameters(flow, source); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_TaskProperty

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{flow, source},
		&returns,
	)

	return returns
}

