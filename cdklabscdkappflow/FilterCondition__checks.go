//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateFilterCondition_BooleanEqualsParameters(field *Field, val interface{}) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case *[]*bool:
		// ok
	case []*bool:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *bool, *[]*bool; received %#v (a %T)", val, val)
	}

	return nil
}

func validateFilterCondition_BooleanNotEqualsParameters(field *Field, val interface{}) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case *[]*bool:
		// ok
	case []*bool:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *bool, *[]*bool; received %#v (a %T)", val, val)
	}

	return nil
}

func validateFilterCondition_NumberEqualsParameters(field *Field, val interface{}) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	case *[]*float64:
		// ok
	case []*float64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *float64, *[]*float64; received %#v (a %T)", val, val)
	}

	return nil
}

func validateFilterCondition_NumberGreaterThanParameters(field *Field, val *float64) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateFilterCondition_NumberGreaterThanEqualsParameters(field *Field, val *float64) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateFilterCondition_NumberLessThanParameters(field *Field, val *float64) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateFilterCondition_NumberLessThanEqualsParameters(field *Field, val *float64) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateFilterCondition_NumberNotEqualsParameters(field *Field, val interface{}) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	case *[]*float64:
		// ok
	case []*float64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *float64, *[]*float64; received %#v (a %T)", val, val)
	}

	return nil
}

func validateFilterCondition_StringContainsParameters(field *Field, val interface{}) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *[]*string:
		// ok
	case []*string:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *[]*string; received %#v (a %T)", val, val)
	}

	return nil
}

func validateFilterCondition_StringEqualsParameters(field *Field, val interface{}) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *[]*string:
		// ok
	case []*string:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *[]*string; received %#v (a %T)", val, val)
	}

	return nil
}

func validateFilterCondition_StringNotEqualsParameters(field *Field, val interface{}) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *[]*string:
		// ok
	case []*string:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *[]*string; received %#v (a %T)", val, val)
	}

	return nil
}

func validateFilterCondition_TimestampBetweenParameters(field *Field, lower *time.Time, upper *time.Time) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if lower == nil {
		return fmt.Errorf("parameter lower is required, but nil was provided")
	}

	if upper == nil {
		return fmt.Errorf("parameter upper is required, but nil was provided")
	}

	return nil
}

func validateFilterCondition_TimestampEqualsParameters(field *Field, val interface{}) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *time.Time:
		// ok
	case time.Time:
		// ok
	case *[]*time.Time:
		// ok
	case []*time.Time:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *time.Time, *[]*time.Time; received %#v (a %T)", val, val)
	}

	return nil
}

func validateFilterCondition_TimestampGreaterThanParameters(field *Field, val interface{}) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *time.Time:
		// ok
	case time.Time:
		// ok
	case *[]*time.Time:
		// ok
	case []*time.Time:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *time.Time, *[]*time.Time; received %#v (a %T)", val, val)
	}

	return nil
}

func validateFilterCondition_TimestampGreaterThanEqualsParameters(field *Field, val interface{}) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *time.Time:
		// ok
	case time.Time:
		// ok
	case *[]*time.Time:
		// ok
	case []*time.Time:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *time.Time, *[]*time.Time; received %#v (a %T)", val, val)
	}

	return nil
}

func validateFilterCondition_TimestampLessThanParameters(field *Field, val interface{}) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *time.Time:
		// ok
	case time.Time:
		// ok
	case *[]*time.Time:
		// ok
	case []*time.Time:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *time.Time, *[]*time.Time; received %#v (a %T)", val, val)
	}

	return nil
}

func validateFilterCondition_TimestampLessThanEqualsParameters(field *Field, val interface{}) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *time.Time:
		// ok
	case time.Time:
		// ok
	case *[]*time.Time:
		// ok
	case []*time.Time:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *time.Time, *[]*time.Time; received %#v (a %T)", val, val)
	}

	return nil
}

func validateFilterCondition_TimestampNotEqualsParameters(field *Field, val interface{}) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *time.Time:
		// ok
	case time.Time:
		// ok
	case *[]*time.Time:
		// ok
	case []*time.Time:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *time.Time, *[]*time.Time; received %#v (a %T)", val, val)
	}

	return nil
}

func validateNewFilterConditionParameters(field *Field, filter *string, properties *TaskProperties) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
		return err
	}

	if filter == nil {
		return fmt.Errorf("parameter filter is required, but nil was provided")
	}

	if properties == nil {
		return fmt.Errorf("parameter properties is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(properties, func() string { return "parameter properties" }); err != nil {
		return err
	}

	return nil
}

