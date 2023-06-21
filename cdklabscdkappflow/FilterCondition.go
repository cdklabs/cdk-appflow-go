package cdklabscdkappflow

import (
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"
)

// A representation of a filter operation condtiion on a source record field.
// Experimental.
type FilterCondition interface {
	// Experimental.
	Field() *Field
	// Experimental.
	Filter() *string
	// Experimental.
	Properties() *TaskProperties
}

// The jsii proxy struct for FilterCondition
type jsiiProxy_FilterCondition struct {
	_ byte // padding
}

func (j *jsiiProxy_FilterCondition) Field() *Field {
	var returns *Field
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilterCondition) Filter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FilterCondition) Properties() *TaskProperties {
	var returns *TaskProperties
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}


// Experimental.
func NewFilterCondition(field *Field, filter *string, properties *TaskProperties) FilterCondition {
	_init_.Initialize()

	if err := validateNewFilterConditionParameters(field, filter, properties); err != nil {
		panic(err)
	}
	j := jsiiProxy_FilterCondition{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.FilterCondition",
		[]interface{}{field, filter, properties},
		&j,
	)

	return &j
}

// Experimental.
func NewFilterCondition_Override(f FilterCondition, field *Field, filter *string, properties *TaskProperties) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.FilterCondition",
		[]interface{}{field, filter, properties},
		f,
	)
}

// Experimental.
func FilterCondition_BooleanEquals(field *Field, val interface{}) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_BooleanEqualsParameters(field, val); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"booleanEquals",
		[]interface{}{field, val},
		&returns,
	)

	return returns
}

// Experimental.
func FilterCondition_BooleanNotEquals(field *Field, val interface{}) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_BooleanNotEqualsParameters(field, val); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"booleanNotEquals",
		[]interface{}{field, val},
		&returns,
	)

	return returns
}

// Experimental.
func FilterCondition_NumberEquals(field *Field, val interface{}) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_NumberEqualsParameters(field, val); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"numberEquals",
		[]interface{}{field, val},
		&returns,
	)

	return returns
}

// Experimental.
func FilterCondition_NumberGreaterThan(field *Field, val *float64) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_NumberGreaterThanParameters(field, val); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"numberGreaterThan",
		[]interface{}{field, val},
		&returns,
	)

	return returns
}

// Experimental.
func FilterCondition_NumberGreaterThanEquals(field *Field, val *float64) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_NumberGreaterThanEqualsParameters(field, val); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"numberGreaterThanEquals",
		[]interface{}{field, val},
		&returns,
	)

	return returns
}

// Experimental.
func FilterCondition_NumberLessThan(field *Field, val *float64) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_NumberLessThanParameters(field, val); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"numberLessThan",
		[]interface{}{field, val},
		&returns,
	)

	return returns
}

// Experimental.
func FilterCondition_NumberLessThanEquals(field *Field, val *float64) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_NumberLessThanEqualsParameters(field, val); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"numberLessThanEquals",
		[]interface{}{field, val},
		&returns,
	)

	return returns
}

// Experimental.
func FilterCondition_NumberNotEquals(field *Field, val interface{}) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_NumberNotEqualsParameters(field, val); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"numberNotEquals",
		[]interface{}{field, val},
		&returns,
	)

	return returns
}

// A condition testing whether a string-type source field contains the given value(s).
//
// NOTE: When multiple values are passed the evaluation is resolved as logical OR.
//
// Returns: an instance of a.
// See: FilterCondition.
//
// Experimental.
func FilterCondition_StringContains(field *Field, val interface{}) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_StringContainsParameters(field, val); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"stringContains",
		[]interface{}{field, val},
		&returns,
	)

	return returns
}

// A condition testing whether a string-type source field equals the given value(s).
//
// NOTE: When multiple values are passed the evaluation is resolved as logical OR.
//
// Returns: an instance of a.
// See: FilterCondition.
//
// Experimental.
func FilterCondition_StringEquals(field *Field, val interface{}) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_StringEqualsParameters(field, val); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"stringEquals",
		[]interface{}{field, val},
		&returns,
	)

	return returns
}

// A condition testing whether a string-type source field does not equal the given value(s).
//
// NOTE: When multiple values are passed the evaluation is resolved as logical OR.
//
// Returns: an instance of a.
// See: FilterCondition.
//
// Experimental.
func FilterCondition_StringNotEquals(field *Field, val interface{}) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_StringNotEqualsParameters(field, val); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"stringNotEquals",
		[]interface{}{field, val},
		&returns,
	)

	return returns
}

// Experimental.
func FilterCondition_TimestampBetween(field *Field, lower *time.Time, upper *time.Time) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_TimestampBetweenParameters(field, lower, upper); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"timestampBetween",
		[]interface{}{field, lower, upper},
		&returns,
	)

	return returns
}

// Experimental.
func FilterCondition_TimestampEquals(field *Field, val interface{}) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_TimestampEqualsParameters(field, val); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"timestampEquals",
		[]interface{}{field, val},
		&returns,
	)

	return returns
}

// Experimental.
func FilterCondition_TimestampGreaterThan(field *Field, val interface{}) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_TimestampGreaterThanParameters(field, val); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"timestampGreaterThan",
		[]interface{}{field, val},
		&returns,
	)

	return returns
}

// Experimental.
func FilterCondition_TimestampGreaterThanEquals(field *Field, val interface{}) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_TimestampGreaterThanEqualsParameters(field, val); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"timestampGreaterThanEquals",
		[]interface{}{field, val},
		&returns,
	)

	return returns
}

// Experimental.
func FilterCondition_TimestampLessThan(field *Field, val interface{}) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_TimestampLessThanParameters(field, val); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"timestampLessThan",
		[]interface{}{field, val},
		&returns,
	)

	return returns
}

// Experimental.
func FilterCondition_TimestampLessThanEquals(field *Field, val interface{}) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_TimestampLessThanEqualsParameters(field, val); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"timestampLessThanEquals",
		[]interface{}{field, val},
		&returns,
	)

	return returns
}

// Experimental.
func FilterCondition_TimestampNotEquals(field *Field, val interface{}) FilterCondition {
	_init_.Initialize()

	if err := validateFilterCondition_TimestampNotEqualsParameters(field, val); err != nil {
		panic(err)
	}
	var returns FilterCondition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.FilterCondition",
		"timestampNotEquals",
		[]interface{}{field, val},
		&returns,
	)

	return returns
}

