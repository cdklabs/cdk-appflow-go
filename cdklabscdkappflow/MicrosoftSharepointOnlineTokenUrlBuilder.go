package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"
)

// A utility class for building Microsoft Online token URLs.
// Experimental.
type MicrosoftSharepointOnlineTokenUrlBuilder interface {
}

// The jsii proxy struct for MicrosoftSharepointOnlineTokenUrlBuilder
type jsiiProxy_MicrosoftSharepointOnlineTokenUrlBuilder struct {
	_ byte // padding
}

// Experimental.
func NewMicrosoftSharepointOnlineTokenUrlBuilder() MicrosoftSharepointOnlineTokenUrlBuilder {
	_init_.Initialize()

	j := jsiiProxy_MicrosoftSharepointOnlineTokenUrlBuilder{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.MicrosoftSharepointOnlineTokenUrlBuilder",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMicrosoftSharepointOnlineTokenUrlBuilder_Override(m MicrosoftSharepointOnlineTokenUrlBuilder) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.MicrosoftSharepointOnlineTokenUrlBuilder",
		nil, // no parameters
		m,
	)
}

// Experimental.
func MicrosoftSharepointOnlineTokenUrlBuilder_BuildTokenUrl(tenantId *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.MicrosoftSharepointOnlineTokenUrlBuilder",
		"buildTokenUrl",
		[]interface{}{tenantId},
		&returns,
	)

	return returns
}

