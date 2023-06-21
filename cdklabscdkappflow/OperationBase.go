package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A base class for all operations.
// Experimental.
type OperationBase interface {
	IOperation
	// Experimental.
	Bind(flow IFlow, source ISource) *[]*awsappflow.CfnFlow_TaskProperty
}

// The jsii proxy struct for OperationBase
type jsiiProxy_OperationBase struct {
	jsiiProxy_IOperation
}

// Experimental.
func NewOperationBase_Override(o OperationBase, tasks *[]ITask) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.OperationBase",
		[]interface{}{tasks},
		o,
	)
}

func (o *jsiiProxy_OperationBase) Bind(flow IFlow, source ISource) *[]*awsappflow.CfnFlow_TaskProperty {
	if err := o.validateBindParameters(flow, source); err != nil {
		panic(err)
	}
	var returns *[]*awsappflow.CfnFlow_TaskProperty

	_jsii_.Invoke(
		o,
		"bind",
		[]interface{}{flow, source},
		&returns,
	)

	return returns
}

