//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HubSpotSource) validateBindParameters(scope IFlow) error {
	return nil
}

func validateNewHubSpotSourceParameters(props *HubSpotSourceProps) error {
	return nil
}

