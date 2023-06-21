package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"
)

// Experimental.
type ValidationAction interface {
	// Experimental.
	Action() *string
}

// The jsii proxy struct for ValidationAction
type jsiiProxy_ValidationAction struct {
	_ byte // padding
}

func (j *jsiiProxy_ValidationAction) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}


// Experimental.
func NewValidationAction(action *string) ValidationAction {
	_init_.Initialize()

	if err := validateNewValidationActionParameters(action); err != nil {
		panic(err)
	}
	j := jsiiProxy_ValidationAction{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.ValidationAction",
		[]interface{}{action},
		&j,
	)

	return &j
}

// Experimental.
func NewValidationAction_Override(v ValidationAction, action *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.ValidationAction",
		[]interface{}{action},
		v,
	)
}

// Returns: a.
// See: ValidationAction that removes a record from the flow execution result.
//
// Experimental.
func ValidationAction_IgnoreRecord() ValidationAction {
	_init_.Initialize()

	var returns ValidationAction

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.ValidationAction",
		"ignoreRecord",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns: a.
// See: ValidationAction that terminates the whole flow execution.
//
// Experimental.
func ValidationAction_TerminateFlow() ValidationAction {
	_init_.Initialize()

	var returns ValidationAction

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.ValidationAction",
		"terminateFlow",
		nil, // no parameters
		&returns,
	)

	return returns
}

