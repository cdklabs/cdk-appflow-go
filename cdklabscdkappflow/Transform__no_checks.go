//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Transform) validateBindParameters(flow IFlow, source ISource) error {
	return nil
}

func validateTransform_MaskParameters(field interface{}) error {
	return nil
}

func validateTransform_MaskEndParameters(field interface{}, length *float64) error {
	return nil
}

func validateTransform_MaskStartParameters(field interface{}, length *float64) error {
	return nil
}

func validateTransform_TruncateParameters(field interface{}, length *float64) error {
	return nil
}

func validateNewTransformParameters(tasks *[]ITask) error {
	return nil
}

