//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"
)

func (f *jsiiProxy_Filter) validateBindParameters(flow IFlow, source ISource) error {
	if flow == nil {
		return fmt.Errorf("parameter flow is required, but nil was provided")
	}

	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}

	return nil
}

func validateFilter_WhenParameters(condition FilterCondition) error {
	if condition == nil {
		return fmt.Errorf("parameter condition is required, but nil was provided")
	}

	return nil
}

func validateNewFilterParameters(condition FilterCondition) error {
	if condition == nil {
		return fmt.Errorf("parameter condition is required, but nil was provided")
	}

	return nil
}

