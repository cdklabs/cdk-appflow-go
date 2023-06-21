//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_SalesforceConnectorProfile) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SalesforceConnectorProfile) validateBuildConnectorProfileCredentialsParameters(properties *ConnectorProfileProps) error {
	if properties == nil {
		return fmt.Errorf("parameter properties is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(properties, func() string { return "parameter properties" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_SalesforceConnectorProfile) validateBuildConnectorProfilePropertiesParameters(properties *ConnectorProfileProps) error {
	if properties == nil {
		return fmt.Errorf("parameter properties is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(properties, func() string { return "parameter properties" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_SalesforceConnectorProfile) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	if arnAttr == nil {
		return fmt.Errorf("parameter arnAttr is required, but nil was provided")
	}

	if arnComponents == nil {
		return fmt.Errorf("parameter arnComponents is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(arnComponents, func() string { return "parameter arnComponents" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_SalesforceConnectorProfile) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SalesforceConnectorProfile) validateTryAddNodeDependencyParameters(scope constructs.IConstruct, resource interface{}) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	switch resource.(type) {
	case *string:
		// ok
	case string:
		// ok
	case constructs.IConstruct:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(resource) {
			return fmt.Errorf("parameter resource must be one of the allowed types: *string, constructs.IConstruct; received %#v (a %T)", resource, resource)
		}
	}

	return nil
}

func validateSalesforceConnectorProfile_FromConnectionProfileArnParameters(scope constructs.Construct, id *string, arn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	return nil
}

func validateSalesforceConnectorProfile_FromConnectionProfileNameParameters(scope constructs.Construct, id *string, name *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateSalesforceConnectorProfile_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateSalesforceConnectorProfile_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateSalesforceConnectorProfile_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewSalesforceConnectorProfileParameters(scope constructs.Construct, id *string, props *SalesforceConnectorProfileProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

