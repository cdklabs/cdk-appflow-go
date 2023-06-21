package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A representation of a validation operation, that is an operation testing records and acting on the test results.
// Experimental.
type Validation interface {
	OperationBase
	IValidation
	// Experimental.
	Action() ValidationAction
	// Experimental.
	Condition() ValidationCondition
	// Experimental.
	Bind(flow IFlow, source ISource) *[]*awsappflow.CfnFlow_TaskProperty
}

// The jsii proxy struct for Validation
type jsiiProxy_Validation struct {
	jsiiProxy_OperationBase
	jsiiProxy_IValidation
}

func (j *jsiiProxy_Validation) Action() ValidationAction {
	var returns ValidationAction
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Validation) Condition() ValidationCondition {
	var returns ValidationCondition
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}


// Experimental.
func NewValidation(condition ValidationCondition, action ValidationAction) Validation {
	_init_.Initialize()

	if err := validateNewValidationParameters(condition, action); err != nil {
		panic(err)
	}
	j := jsiiProxy_Validation{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.Validation",
		[]interface{}{condition, action},
		&j,
	)

	return &j
}

// Experimental.
func NewValidation_Override(v Validation, condition ValidationCondition, action ValidationAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.Validation",
		[]interface{}{condition, action},
		v,
	)
}

// Returns: a Validation instance.
// See: ValidationAction for the validation.
//
// Experimental.
func Validation_When(condition ValidationCondition, action ValidationAction) IValidation {
	_init_.Initialize()

	if err := validateValidation_WhenParameters(condition, action); err != nil {
		panic(err)
	}
	var returns IValidation

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.Validation",
		"when",
		[]interface{}{condition, action},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Validation) Bind(flow IFlow, source ISource) *[]*awsappflow.CfnFlow_TaskProperty {
	if err := v.validateBindParameters(flow, source); err != nil {
		panic(err)
	}
	var returns *[]*awsappflow.CfnFlow_TaskProperty

	_jsii_.Invoke(
		v,
		"bind",
		[]interface{}{flow, source},
		&returns,
	)

	return returns
}

