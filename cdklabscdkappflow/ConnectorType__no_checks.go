//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func validateNewConnectorTypeParameters(name *string, isCustom *bool) error {
	return nil
}

