//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RedshiftDestination) validateBindParameters(scope IFlow) error {
	return nil
}

func validateNewRedshiftDestinationParameters(props *RedshiftDestinationProps) error {
	return nil
}

