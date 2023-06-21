//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"
)

func validateNewWriteOperationParameters(type_ WriteOperationType) error {
	if type_ == "" {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

