//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"
)

func (o *jsiiProxy_OperationBase) validateBindParameters(flow IFlow, source ISource) error {
	if flow == nil {
		return fmt.Errorf("parameter flow is required, but nil was provided")
	}

	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}

	return nil
}

func validateNewOperationBaseParameters(tasks *[]ITask) error {
	if tasks == nil {
		return fmt.Errorf("parameter tasks is required, but nil was provided")
	}

	return nil
}

