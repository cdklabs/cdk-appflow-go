//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"
)

func validateMicrosoftDynamics365ApiUrlBuilder_BuildApiUrlParameters(org *string) error {
	if org == nil {
		return fmt.Errorf("parameter org is required, but nil was provided")
	}

	return nil
}

