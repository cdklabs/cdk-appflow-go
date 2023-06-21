//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TriggeredFlowBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TriggeredFlowBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TriggeredFlowBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (t *jsiiProxy_TriggeredFlowBase) validateOnDeactivatedParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (t *jsiiProxy_TriggeredFlowBase) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (t *jsiiProxy_TriggeredFlowBase) validateOnRunCompletedParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (t *jsiiProxy_TriggeredFlowBase) validateOnRunStartedParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateTriggeredFlowBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTriggeredFlowBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTriggeredFlowBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTriggeredFlowBaseParameters(scope constructs.Construct, id *string, props *FlowBaseProps) error {
	return nil
}

