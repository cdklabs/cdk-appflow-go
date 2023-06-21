//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"
)

func validateServiceNowInstanceUrlBuilder_BuildFromDomainParameters(domain *string) error {
	if domain == nil {
		return fmt.Errorf("parameter domain is required, but nil was provided")
	}

	return nil
}

