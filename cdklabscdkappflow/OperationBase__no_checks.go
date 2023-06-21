//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OperationBase) validateBindParameters(flow IFlow, source ISource) error {
	return nil
}

func validateNewOperationBaseParameters(tasks *[]ITask) error {
	return nil
}

