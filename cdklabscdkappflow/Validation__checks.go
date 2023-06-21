//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"
)

func (v *jsiiProxy_Validation) validateBindParameters(flow IFlow, source ISource) error {
	if flow == nil {
		return fmt.Errorf("parameter flow is required, but nil was provided")
	}

	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}

	return nil
}

func validateValidation_WhenParameters(condition ValidationCondition, action ValidationAction) error {
	if condition == nil {
		return fmt.Errorf("parameter condition is required, but nil was provided")
	}

	if action == nil {
		return fmt.Errorf("parameter action is required, but nil was provided")
	}

	return nil
}

func validateNewValidationParameters(condition ValidationCondition, action ValidationAction) error {
	if condition == nil {
		return fmt.Errorf("parameter condition is required, but nil was provided")
	}

	if action == nil {
		return fmt.Errorf("parameter action is required, but nil was provided")
	}

	return nil
}

