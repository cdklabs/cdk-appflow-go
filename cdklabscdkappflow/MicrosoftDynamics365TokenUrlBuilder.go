package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"
)

// A utility class for building Microsoft Dynamics 365 token URLs.
// Experimental.
type MicrosoftDynamics365TokenUrlBuilder interface {
}

// The jsii proxy struct for MicrosoftDynamics365TokenUrlBuilder
type jsiiProxy_MicrosoftDynamics365TokenUrlBuilder struct {
	_ byte // padding
}

// Experimental.
func NewMicrosoftDynamics365TokenUrlBuilder() MicrosoftDynamics365TokenUrlBuilder {
	_init_.Initialize()

	j := jsiiProxy_MicrosoftDynamics365TokenUrlBuilder{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365TokenUrlBuilder",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMicrosoftDynamics365TokenUrlBuilder_Override(m MicrosoftDynamics365TokenUrlBuilder) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365TokenUrlBuilder",
		nil, // no parameters
		m,
	)
}

// Experimental.
func MicrosoftDynamics365TokenUrlBuilder_BuildTokenUrl(tenantId *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365TokenUrlBuilder",
		"buildTokenUrl",
		[]interface{}{tenantId},
		&returns,
	)

	return returns
}

