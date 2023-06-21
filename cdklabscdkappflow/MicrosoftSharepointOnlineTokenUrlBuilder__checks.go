//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"
)

func validateMicrosoftSharepointOnlineTokenUrlBuilder_BuildFromTenantParameters(tenantId *string) error {
	if tenantId == nil {
		return fmt.Errorf("parameter tenantId is required, but nil was provided")
	}

	return nil
}

