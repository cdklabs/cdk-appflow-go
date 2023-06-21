//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3Source) validateBindParameters(scope IFlow) error {
	return nil
}

func validateNewS3SourceParameters(props *S3SourceProps) error {
	return nil
}

