//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) validateBuildConnectorProfileCredentialsParameters(props *ConnectorProfileProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) validateBuildConnectorProfilePropertiesParameters(_props *ConnectorProfileProps) error {
	if _props == nil {
		return fmt.Errorf("parameter _props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_props, func() string { return "parameter _props" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
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

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JdbcSmallDataScaleConnectorProfile) validateTryAddNodeDependencyParameters(scope constructs.IConstruct, resource interface{}) error {
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

func validateJdbcSmallDataScaleConnectorProfile_FromConnectionProfileArnParameters(scope constructs.Construct, id *string, arn *string) error {
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

func validateJdbcSmallDataScaleConnectorProfile_FromConnectionProfileNameParameters(scope constructs.Construct, id *string, name *string) error {
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

func validateJdbcSmallDataScaleConnectorProfile_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateJdbcSmallDataScaleConnectorProfile_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateJdbcSmallDataScaleConnectorProfile_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewJdbcSmallDataScaleConnectorProfileParameters(scope constructs.Construct, id *string, props *JdbcSmallDataScaleConnectorProfileProps) error {
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

