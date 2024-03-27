//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MailchimpSource) validateBindParameters(flow IFlow) error {
	return nil
}

func validateNewMailchimpSourceParameters(props *MailchimpSourceProps) error {
	return nil
}

