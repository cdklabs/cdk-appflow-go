//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Task) validateBindParameters(_flow IFlow, source ISource) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetConnectorOperatorParameters(val *TaskConnectorOperator) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetPropertiesParameters(val *TaskProperties) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetSourceFieldsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewTaskParameters(type_ *string, sourceFields *[]*string, connectorOperator *TaskConnectorOperator, properties *TaskProperties) error {
	return nil
}

