//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SAPOdataDestination) validateBindParameters(flow IFlow) error {
	return nil
}

func validateNewSAPOdataDestinationParameters(props *SAPOdataDestinationProps) error {
	return nil
}

