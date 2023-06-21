//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_Mapping) validateBindParameters(flow IFlow, source ISource) error {
	return nil
}

func validateMapping_AddParameters(sourceField1 *Field, sourceField2 *Field, to *Field) error {
	return nil
}

func validateMapping_ConcatParameters(from *[]*Field, to *Field, format *string) error {
	return nil
}

func validateMapping_DivideParameters(sourceField1 *Field, sourceField2 *Field, to *Field) error {
	return nil
}

func validateMapping_MapParameters(from *Field, to *Field) error {
	return nil
}

func validateMapping_MapAllParameters(config *MapAllConfig) error {
	return nil
}

func validateMapping_MultiplyParameters(sourceField1 *Field, sourceField2 *Field, to *Field) error {
	return nil
}

func validateMapping_SubtractParameters(sourceField1 *Field, sourceField2 *Field, to *Field) error {
	return nil
}

func validateNewMappingParameters(tasks *[]ITask) error {
	return nil
}

