//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnowflakeDestination) validateBindParameters(scope IFlow) error {
	return nil
}

func validateNewSnowflakeDestinationParameters(props *SnowflakeDestinationProps) error {
	return nil
}

