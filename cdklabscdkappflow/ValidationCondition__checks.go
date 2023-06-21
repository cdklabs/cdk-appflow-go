//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateValidationCondition_IsDefaultParameters(field interface{}) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	switch field.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *Field:
		field := field.(*Field)
		if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
			return err
		}
	case Field:
		field_ := field.(Field)
		field := &field_
		if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(field) {
			return fmt.Errorf("parameter field must be one of the allowed types: *string, *Field; received %#v (a %T)", field, field)
		}
	}

	return nil
}

func validateValidationCondition_IsNegativeParameters(field interface{}) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	switch field.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *Field:
		field := field.(*Field)
		if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
			return err
		}
	case Field:
		field_ := field.(Field)
		field := &field_
		if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(field) {
			return fmt.Errorf("parameter field must be one of the allowed types: *string, *Field; received %#v (a %T)", field, field)
		}
	}

	return nil
}

func validateValidationCondition_IsNotNullParameters(field interface{}) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	switch field.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *Field:
		field := field.(*Field)
		if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
			return err
		}
	case Field:
		field_ := field.(Field)
		field := &field_
		if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(field) {
			return fmt.Errorf("parameter field must be one of the allowed types: *string, *Field; received %#v (a %T)", field, field)
		}
	}

	return nil
}

func validateValidationCondition_IsNullParameters(field interface{}) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}
	switch field.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *Field:
		field := field.(*Field)
		if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
			return err
		}
	case Field:
		field_ := field.(Field)
		field := &field_
		if err := _jsii_.ValidateStruct(field, func() string { return "parameter field" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(field) {
			return fmt.Errorf("parameter field must be one of the allowed types: *string, *Field; received %#v (a %T)", field, field)
		}
	}

	return nil
}

func validateNewValidationConditionParameters(field *string, validation *string) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}

	if validation == nil {
		return fmt.Errorf("parameter validation is required, but nil was provided")
	}

	return nil
}

