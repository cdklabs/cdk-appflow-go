//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_Filter) validateBindParameters(flow IFlow, source ISource) error {
	return nil
}

func validateFilter_WhenParameters(condition FilterCondition) error {
	return nil
}

func validateNewFilterParameters(condition FilterCondition) error {
	return nil
}

