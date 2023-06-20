package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"
)

// Experimental.
type Hello interface {
	// Experimental.
	SayHello() *string
}

// The jsii proxy struct for Hello
type jsiiProxy_Hello struct {
	_ byte // padding
}

// Experimental.
func NewHello() Hello {
	_init_.Initialize()

	j := jsiiProxy_Hello{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.Hello",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewHello_Override(h Hello) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.Hello",
		nil, // no parameters
		h,
	)
}

func (h *jsiiProxy_Hello) SayHello() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"sayHello",
		nil, // no parameters
		&returns,
	)

	return returns
}

