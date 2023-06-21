//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"
)

func validateNewConnectorTypeParameters(name *string, isCustom *bool) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if isCustom == nil {
		return fmt.Errorf("parameter isCustom is required, but nil was provided")
	}

	return nil
}

