//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"
)

func validateNewValidationActionParameters(action *string) error {
	if action == nil {
		return fmt.Errorf("parameter action is required, but nil was provided")
	}

	return nil
}

