//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IFlow) validateOnRunCompletedParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IFlow) validateOnRunStartedParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

