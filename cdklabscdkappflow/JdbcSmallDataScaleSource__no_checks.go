//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JdbcSmallDataScaleSource) validateBindParameters(flow IFlow) error {
	return nil
}

func validateNewJdbcSmallDataScaleSourceParameters(props *JdbcSmallDataScaleSourceProps) error {
	return nil
}

