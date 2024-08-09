//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AsanaSource) validateBindParameters(flow IFlow) error {
	return nil
}

func validateNewAsanaSourceParameters(props *AsanaSourceProps) error {
	return nil
}

