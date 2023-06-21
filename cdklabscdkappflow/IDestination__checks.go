//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"
)

func (i *jsiiProxy_IDestination) validateBindParameters(scope IFlow) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

