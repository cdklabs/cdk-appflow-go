package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"
)

// A representation of a validation condition on a particular field in a flow execution.
// Experimental.
type ValidationCondition interface {
	// Experimental.
	Field() *string
	// Experimental.
	Validation() *string
}

// The jsii proxy struct for ValidationCondition
type jsiiProxy_ValidationCondition struct {
	_ byte // padding
}

func (j *jsiiProxy_ValidationCondition) Field() *string {
	var returns *string
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidationCondition) Validation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validation",
		&returns,
	)
	return returns
}


// Experimental.
func NewValidationCondition(field *string, validation *string) ValidationCondition {
	_init_.Initialize()

	if err := validateNewValidationConditionParameters(field, validation); err != nil {
		panic(err)
	}
	j := jsiiProxy_ValidationCondition{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.ValidationCondition",
		[]interface{}{field, validation},
		&j,
	)

	return &j
}

// Experimental.
func NewValidationCondition_Override(v ValidationCondition, field *string, validation *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.ValidationCondition",
		[]interface{}{field, validation},
		v,
	)
}

// Returns: a.
// See: ValidationCondition instance.
//
// Experimental.
func ValidationCondition_IsDefault(field interface{}) ValidationCondition {
	_init_.Initialize()

	if err := validateValidationCondition_IsDefaultParameters(field); err != nil {
		panic(err)
	}
	var returns ValidationCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.ValidationCondition",
		"isDefault",
		[]interface{}{field},
		&returns,
	)

	return returns
}

// Validates whether a particular field in an execution is negative.
//
// Returns: a.
// See: ValidationCondition instance.
//
// Experimental.
func ValidationCondition_IsNegative(field interface{}) ValidationCondition {
	_init_.Initialize()

	if err := validateValidationCondition_IsNegativeParameters(field); err != nil {
		panic(err)
	}
	var returns ValidationCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.ValidationCondition",
		"isNegative",
		[]interface{}{field},
		&returns,
	)

	return returns
}

// Returns: a.
// See: ValidationCondition instance.
//
// Experimental.
func ValidationCondition_IsNotNull(field interface{}) ValidationCondition {
	_init_.Initialize()

	if err := validateValidationCondition_IsNotNullParameters(field); err != nil {
		panic(err)
	}
	var returns ValidationCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.ValidationCondition",
		"isNotNull",
		[]interface{}{field},
		&returns,
	)

	return returns
}

// Validates whether a particular field has no value.
//
// Returns: a.
// See: ValidationCondition instance.
//
// Experimental.
func ValidationCondition_IsNull(field interface{}) ValidationCondition {
	_init_.Initialize()

	if err := validateValidationCondition_IsNullParameters(field); err != nil {
		panic(err)
	}
	var returns ValidationCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.ValidationCondition",
		"isNull",
		[]interface{}{field},
		&returns,
	)

	return returns
}

