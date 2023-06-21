package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"
)

// Experimental.
type ZendeskInstanceUrlBuilder interface {
}

// The jsii proxy struct for ZendeskInstanceUrlBuilder
type jsiiProxy_ZendeskInstanceUrlBuilder struct {
	_ byte // padding
}

// Experimental.
func NewZendeskInstanceUrlBuilder() ZendeskInstanceUrlBuilder {
	_init_.Initialize()

	j := jsiiProxy_ZendeskInstanceUrlBuilder{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.ZendeskInstanceUrlBuilder",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewZendeskInstanceUrlBuilder_Override(z ZendeskInstanceUrlBuilder) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.ZendeskInstanceUrlBuilder",
		nil, // no parameters
		z,
	)
}

// Experimental.
func ZendeskInstanceUrlBuilder_BuildFromAccount(account *string) *string {
	_init_.Initialize()

	if err := validateZendeskInstanceUrlBuilder_BuildFromAccountParameters(account); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.ZendeskInstanceUrlBuilder",
		"buildFromAccount",
		[]interface{}{account},
		&returns,
	)

	return returns
}

