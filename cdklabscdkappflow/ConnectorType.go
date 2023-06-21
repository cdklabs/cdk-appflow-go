package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"
)

// Experimental.
type ConnectorType interface {
	// Experimental.
	AsProfileConnectorLabel() *string
	// Experimental.
	AsProfileConnectorType() *string
	// Experimental.
	AsTaskConnectorOperatorOrigin() *string
	// Experimental.
	IsCustom() *bool
	// Experimental.
	Name() *string
}

// The jsii proxy struct for ConnectorType
type jsiiProxy_ConnectorType struct {
	_ byte // padding
}

func (j *jsiiProxy_ConnectorType) AsProfileConnectorLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asProfileConnectorLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectorType) AsProfileConnectorType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asProfileConnectorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectorType) AsTaskConnectorOperatorOrigin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asTaskConnectorOperatorOrigin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectorType) IsCustom() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isCustom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectorType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewConnectorType(name *string, isCustom *bool) ConnectorType {
	_init_.Initialize()

	if err := validateNewConnectorTypeParameters(name, isCustom); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectorType{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.ConnectorType",
		[]interface{}{name, isCustom},
		&j,
	)

	return &j
}

// Experimental.
func NewConnectorType_Override(c ConnectorType, name *string, isCustom *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.ConnectorType",
		[]interface{}{name, isCustom},
		c,
	)
}

