//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (z *jsiiProxy_ZendeskSource) validateBindParameters(flow IFlow) error {
	return nil
}

func validateNewZendeskSourceParameters(props *ZendeskSourceProps) error {
	return nil
}

