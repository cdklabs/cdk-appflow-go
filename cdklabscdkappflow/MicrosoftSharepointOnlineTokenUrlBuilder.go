package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"
)

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
func MicrosoftSharepointOnlineTokenUrlBuilder_BuildFromTenant(tenantId *string) *string {
	_init_.Initialize()

	if err := validateMicrosoftSharepointOnlineTokenUrlBuilder_BuildFromTenantParameters(tenantId); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.MicrosoftSharepointOnlineTokenUrlBuilder",
		"buildFromTenant",
		[]interface{}{tenantId},
		&returns,
	)

	return returns
}

