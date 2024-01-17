//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GoogleBigQuerySource) validateBindParameters(scope IFlow) error {
	return nil
}

func validateNewGoogleBigQuerySourceParameters(props *GoogleBigQuerySourceProps) error {
	return nil
}

