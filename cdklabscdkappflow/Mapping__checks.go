//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (m *jsiiProxy_Mapping) validateBindParameters(flow IFlow, source ISource) error {
	if flow == nil {
		return fmt.Errorf("parameter flow is required, but nil was provided")
	}

	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}

	return nil
}

func validateMapping_AddParameters(sourceField1 *Field, sourceField2 *Field, to *Field) error {
	if sourceField1 == nil {
		return fmt.Errorf("parameter sourceField1 is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(sourceField1, func() string { return "parameter sourceField1" }); err != nil {
		return err
	}

	if sourceField2 == nil {
		return fmt.Errorf("parameter sourceField2 is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(sourceField2, func() string { return "parameter sourceField2" }); err != nil {
		return err
	}

	if to == nil {
		return fmt.Errorf("parameter to is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(to, func() string { return "parameter to" }); err != nil {
		return err
	}

	return nil
}

func validateMapping_ConcatParameters(from *[]*Field, to *Field, format *string) error {
	if from == nil {
		return fmt.Errorf("parameter from is required, but nil was provided")
	}
	for idx_75857a, v := range *from {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter from[%#v]", idx_75857a) }); err != nil {
			return err
		}
	}

	if to == nil {
		return fmt.Errorf("parameter to is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(to, func() string { return "parameter to" }); err != nil {
		return err
	}

	if format == nil {
		return fmt.Errorf("parameter format is required, but nil was provided")
	}

	return nil
}

func validateMapping_DivideParameters(sourceField1 *Field, sourceField2 *Field, to *Field) error {
	if sourceField1 == nil {
		return fmt.Errorf("parameter sourceField1 is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(sourceField1, func() string { return "parameter sourceField1" }); err != nil {
		return err
	}

	if sourceField2 == nil {
		return fmt.Errorf("parameter sourceField2 is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(sourceField2, func() string { return "parameter sourceField2" }); err != nil {
		return err
	}

	if to == nil {
		return fmt.Errorf("parameter to is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(to, func() string { return "parameter to" }); err != nil {
		return err
	}

	return nil
}

func validateMapping_MapParameters(from *Field, to *Field) error {
	if from == nil {
		return fmt.Errorf("parameter from is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(from, func() string { return "parameter from" }); err != nil {
		return err
	}

	if to == nil {
		return fmt.Errorf("parameter to is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(to, func() string { return "parameter to" }); err != nil {
		return err
	}

	return nil
}

func validateMapping_MapAllParameters(config *MapAllConfig) error {
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

func validateMapping_MultiplyParameters(sourceField1 *Field, sourceField2 *Field, to *Field) error {
	if sourceField1 == nil {
		return fmt.Errorf("parameter sourceField1 is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(sourceField1, func() string { return "parameter sourceField1" }); err != nil {
		return err
	}

	if sourceField2 == nil {
		return fmt.Errorf("parameter sourceField2 is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(sourceField2, func() string { return "parameter sourceField2" }); err != nil {
		return err
	}

	if to == nil {
		return fmt.Errorf("parameter to is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(to, func() string { return "parameter to" }); err != nil {
		return err
	}

	return nil
}

func validateMapping_SubtractParameters(sourceField1 *Field, sourceField2 *Field, to *Field) error {
	if sourceField1 == nil {
		return fmt.Errorf("parameter sourceField1 is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(sourceField1, func() string { return "parameter sourceField1" }); err != nil {
		return err
	}

	if sourceField2 == nil {
		return fmt.Errorf("parameter sourceField2 is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(sourceField2, func() string { return "parameter sourceField2" }); err != nil {
		return err
	}

	if to == nil {
		return fmt.Errorf("parameter to is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(to, func() string { return "parameter to" }); err != nil {
		return err
	}

	return nil
}

func validateNewMappingParameters(tasks *[]ITask) error {
	if tasks == nil {
		return fmt.Errorf("parameter tasks is required, but nil was provided")
	}

	return nil
}

