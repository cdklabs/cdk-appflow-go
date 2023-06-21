package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"
)

// Experimental.
type ServiceNowInstanceUrlBuilder interface {
}

// The jsii proxy struct for ServiceNowInstanceUrlBuilder
type jsiiProxy_ServiceNowInstanceUrlBuilder struct {
	_ byte // padding
}

// Experimental.
func NewServiceNowInstanceUrlBuilder() ServiceNowInstanceUrlBuilder {
	_init_.Initialize()

	j := jsiiProxy_ServiceNowInstanceUrlBuilder{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.ServiceNowInstanceUrlBuilder",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewServiceNowInstanceUrlBuilder_Override(s ServiceNowInstanceUrlBuilder) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.ServiceNowInstanceUrlBuilder",
		nil, // no parameters
		s,
	)
}

// Experimental.
func ServiceNowInstanceUrlBuilder_BuildFromDomain(domain *string) *string {
	_init_.Initialize()

	if err := validateServiceNowInstanceUrlBuilder_BuildFromDomainParameters(domain); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.ServiceNowInstanceUrlBuilder",
		"buildFromDomain",
		[]interface{}{domain},
		&returns,
	)

	return returns
}

