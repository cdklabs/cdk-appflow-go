package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A representation of a transform operation, that is an operation modifying source fields.
// Experimental.
type Transform interface {
	OperationBase
	ITransform
	// Experimental.
	Bind(flow IFlow, source ISource) *[]*awsappflow.CfnFlow_TaskProperty
}

// The jsii proxy struct for Transform
type jsiiProxy_Transform struct {
	jsiiProxy_OperationBase
	jsiiProxy_ITransform
}

// Experimental.
func NewTransform(tasks *[]ITask) Transform {
	_init_.Initialize()

	if err := validateNewTransformParameters(tasks); err != nil {
		panic(err)
	}
	j := jsiiProxy_Transform{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.Transform",
		[]interface{}{tasks},
		&j,
	)

	return &j
}

// Experimental.
func NewTransform_Override(t Transform, tasks *[]ITask) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.Transform",
		[]interface{}{tasks},
		t,
	)
}

// Masks the field with a specified mask.
//
// Returns: a.
// See: Transform instance.
//
// Experimental.
func Transform_Mask(field interface{}, mask *string) ITransform {
	_init_.Initialize()

	if err := validateTransform_MaskParameters(field); err != nil {
		panic(err)
	}
	var returns ITransform

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.Transform",
		"mask",
		[]interface{}{field, mask},
		&returns,
	)

	return returns
}

// Experimental.
func Transform_MaskEnd(field interface{}, length *float64, mask *string) ITransform {
	_init_.Initialize()

	if err := validateTransform_MaskEndParameters(field, length); err != nil {
		panic(err)
	}
	var returns ITransform

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.Transform",
		"maskEnd",
		[]interface{}{field, length, mask},
		&returns,
	)

	return returns
}

// Experimental.
func Transform_MaskStart(field interface{}, length *float64, mask *string) ITransform {
	_init_.Initialize()

	if err := validateTransform_MaskStartParameters(field, length); err != nil {
		panic(err)
	}
	var returns ITransform

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.Transform",
		"maskStart",
		[]interface{}{field, length, mask},
		&returns,
	)

	return returns
}

// Truncates the field to a specified length.
//
// Returns: a.
// See: Transform instance.
//
// Experimental.
func Transform_Truncate(field interface{}, length *float64) ITransform {
	_init_.Initialize()

	if err := validateTransform_TruncateParameters(field, length); err != nil {
		panic(err)
	}
	var returns ITransform

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.Transform",
		"truncate",
		[]interface{}{field, length},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Transform) Bind(flow IFlow, source ISource) *[]*awsappflow.CfnFlow_TaskProperty {
	if err := t.validateBindParameters(flow, source); err != nil {
		panic(err)
	}
	var returns *[]*awsappflow.CfnFlow_TaskProperty

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{flow, source},
		&returns,
	)

	return returns
}

