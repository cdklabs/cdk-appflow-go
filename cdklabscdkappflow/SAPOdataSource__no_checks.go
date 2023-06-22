//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SAPOdataSource) validateBindParameters(flow IFlow) error {
	return nil
}

func validateNewSAPOdataSourceParameters(props *SAPOdataSourceProps) error {
	return nil
}

