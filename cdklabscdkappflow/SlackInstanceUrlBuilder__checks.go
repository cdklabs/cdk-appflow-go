//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"
)

func validateSlackInstanceUrlBuilder_BuildFromWorkspaceParameters(workspace *string) error {
	if workspace == nil {
		return fmt.Errorf("parameter workspace is required, but nil was provided")
	}

	return nil
}

