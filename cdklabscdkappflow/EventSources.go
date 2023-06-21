package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"
)

// Experimental.
type EventSources interface {
}

// The jsii proxy struct for EventSources
type jsiiProxy_EventSources struct {
	_ byte // padding
}

// Experimental.
func NewEventSources() EventSources {
	_init_.Initialize()

	j := jsiiProxy_EventSources{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.EventSources",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEventSources_Override(e EventSources) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.EventSources",
		nil, // no parameters
		e,
	)
}

// Experimental.
func EventSources_SalesforceEventSource(suffix *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.EventSources",
		"salesforceEventSource",
		[]interface{}{suffix},
		&returns,
	)

	return returns
}

