//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (t *jsiiProxy_Transform) validateBindParameters(flow IFlow, source ISource) error {
	if flow == nil {
		return fmt.Errorf("parameter flow is required, but nil was provided")
	}

	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}

	return nil
}

func validateTransform_MaskParameters(field interface{}) error {
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

func validateTransform_MaskEndParameters(field interface{}, length *float64) error {
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

	if length == nil {
		return fmt.Errorf("parameter length is required, but nil was provided")
	}

	return nil
}

func validateTransform_MaskStartParameters(field interface{}, length *float64) error {
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

	if length == nil {
		return fmt.Errorf("parameter length is required, but nil was provided")
	}

	return nil
}

func validateTransform_TruncateParameters(field interface{}, length *float64) error {
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

	if length == nil {
		return fmt.Errorf("parameter length is required, but nil was provided")
	}

	return nil
}

func validateNewTransformParameters(tasks *[]ITask) error {
	if tasks == nil {
		return fmt.Errorf("parameter tasks is required, but nil was provided")
	}

	return nil
}

