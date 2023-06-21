//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"
)

func validateZendeskInstanceUrlBuilder_BuildFromAccountParameters(account *string) error {
	if account == nil {
		return fmt.Errorf("parameter account is required, but nil was provided")
	}

	return nil
}

