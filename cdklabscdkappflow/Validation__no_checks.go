//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_Validation) validateBindParameters(flow IFlow, source ISource) error {
	return nil
}

func validateValidation_WhenParameters(condition ValidationCondition, action ValidationAction) error {
	return nil
}

func validateNewValidationParameters(condition ValidationCondition, action ValidationAction) error {
	return nil
}

