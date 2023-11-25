package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"
)

// Experimental.
type MicrosoftDynamics365ApiUrlBuilder interface {
}

// The jsii proxy struct for MicrosoftDynamics365ApiUrlBuilder
type jsiiProxy_MicrosoftDynamics365ApiUrlBuilder struct {
	_ byte // padding
}

// Experimental.
func NewMicrosoftDynamics365ApiUrlBuilder() MicrosoftDynamics365ApiUrlBuilder {
	_init_.Initialize()

	j := jsiiProxy_MicrosoftDynamics365ApiUrlBuilder{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365ApiUrlBuilder",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMicrosoftDynamics365ApiUrlBuilder_Override(m MicrosoftDynamics365ApiUrlBuilder) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365ApiUrlBuilder",
		nil, // no parameters
		m,
	)
}

// Experimental.
func MicrosoftDynamics365ApiUrlBuilder_BuildApiUrl(org *string) *string {
	_init_.Initialize()

	if err := validateMicrosoftDynamics365ApiUrlBuilder_BuildApiUrlParameters(org); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365ApiUrlBuilder",
		"buildApiUrl",
		[]interface{}{org},
		&returns,
	)

	return returns
}

