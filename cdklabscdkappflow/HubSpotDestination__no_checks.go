//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HubSpotDestination) validateBindParameters(flow IFlow) error {
	return nil
}

func validateNewHubSpotDestinationParameters(props *HubSpotDestinationProps) error {
	return nil
}

