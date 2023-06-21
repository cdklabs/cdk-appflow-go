package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"
)

// Experimental.
type SlackInstanceUrlBuilder interface {
}

// The jsii proxy struct for SlackInstanceUrlBuilder
type jsiiProxy_SlackInstanceUrlBuilder struct {
	_ byte // padding
}

// Experimental.
func NewSlackInstanceUrlBuilder() SlackInstanceUrlBuilder {
	_init_.Initialize()

	j := jsiiProxy_SlackInstanceUrlBuilder{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.SlackInstanceUrlBuilder",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSlackInstanceUrlBuilder_Override(s SlackInstanceUrlBuilder) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.SlackInstanceUrlBuilder",
		nil, // no parameters
		s,
	)
}

// Experimental.
func SlackInstanceUrlBuilder_BuildFromWorkspace(workspace *string) *string {
	_init_.Initialize()

	if err := validateSlackInstanceUrlBuilder_BuildFromWorkspaceParameters(workspace); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.SlackInstanceUrlBuilder",
		"buildFromWorkspace",
		[]interface{}{workspace},
		&returns,
	)

	return returns
}

