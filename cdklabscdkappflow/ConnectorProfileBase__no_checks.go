//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConnectorProfileBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ConnectorProfileBase) validateBuildConnectorProfileCredentialsParameters(props *ConnectorProfileProps) error {
	return nil
}

func (c *jsiiProxy_ConnectorProfileBase) validateBuildConnectorProfilePropertiesParameters(props *ConnectorProfileProps) error {
	return nil
}

func (c *jsiiProxy_ConnectorProfileBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ConnectorProfileBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_ConnectorProfileBase) validateTryAddNodeDependencyParameters(scope constructs.IConstruct, resource interface{}) error {
	return nil
}

func validateConnectorProfileBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateConnectorProfileBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateConnectorProfileBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewConnectorProfileBaseParameters(scope constructs.Construct, id *string, props *ConnectorProfileProps, connectorType ConnectorType) error {
	return nil
}

