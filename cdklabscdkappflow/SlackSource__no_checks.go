//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SlackSource) validateBindParameters(flow IFlow) error {
	return nil
}

func validateNewSlackSourceParameters(props *SlackSourceProps) error {
	return nil
}

