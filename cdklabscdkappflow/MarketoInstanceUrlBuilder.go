package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"
)

// Experimental.
type MarketoInstanceUrlBuilder interface {
}

// The jsii proxy struct for MarketoInstanceUrlBuilder
type jsiiProxy_MarketoInstanceUrlBuilder struct {
	_ byte // padding
}

// Experimental.
func NewMarketoInstanceUrlBuilder() MarketoInstanceUrlBuilder {
	_init_.Initialize()

	j := jsiiProxy_MarketoInstanceUrlBuilder{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.MarketoInstanceUrlBuilder",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMarketoInstanceUrlBuilder_Override(m MarketoInstanceUrlBuilder) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.MarketoInstanceUrlBuilder",
		nil, // no parameters
		m,
	)
}

// Experimental.
func MarketoInstanceUrlBuilder_BuildFromAccount(account *string) *string {
	_init_.Initialize()

	if err := validateMarketoInstanceUrlBuilder_BuildFromAccountParameters(account); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.MarketoInstanceUrlBuilder",
		"buildFromAccount",
		[]interface{}{account},
		&returns,
	)

	return returns
}

