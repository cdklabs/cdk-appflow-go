//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (t *jsiiProxy_Task) validateBindParameters(_flow IFlow, source ISource) error {
	if _flow == nil {
		return fmt.Errorf("parameter _flow is required, but nil was provided")
	}

	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Task) validateSetConnectorOperatorParameters(val *TaskConnectorOperator) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_Task) validateSetPropertiesParameters(val *[]*TaskProperty) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	for idx_97dfc6, v := range *val {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
			return err
		}
	}

	return nil
}

func (j *jsiiProxy_Task) validateSetSourceFieldsParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Task) validateSetTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTaskParameters(type_ *string, sourceFields *[]*string, connectorOperator *TaskConnectorOperator, properties *[]*TaskProperty) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	if sourceFields == nil {
		return fmt.Errorf("parameter sourceFields is required, but nil was provided")
	}

	if connectorOperator == nil {
		return fmt.Errorf("parameter connectorOperator is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(connectorOperator, func() string { return "parameter connectorOperator" }); err != nil {
		return err
	}

	if properties == nil {
		return fmt.Errorf("parameter properties is required, but nil was provided")
	}
	for idx_96adb9, v := range *properties {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter properties[%#v]", idx_96adb9) }); err != nil {
			return err
		}
	}

	return nil
}

