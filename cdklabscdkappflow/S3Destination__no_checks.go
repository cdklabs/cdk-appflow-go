//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3Destination) validateBindParameters(flow IFlow) error {
	return nil
}

func validateNewS3DestinationParameters(props *S3DestinationProps) error {
	return nil
}

