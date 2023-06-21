package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A representation of an instance of a mapping operation, that is an operation translating source to destination fields.
// Experimental.
type Mapping interface {
	OperationBase
	IMapping
	// Experimental.
	Bind(flow IFlow, source ISource) *[]*awsappflow.CfnFlow_TaskProperty
}

// The jsii proxy struct for Mapping
type jsiiProxy_Mapping struct {
	jsiiProxy_OperationBase
	jsiiProxy_IMapping
}

// Experimental.
func NewMapping(tasks *[]ITask) Mapping {
	_init_.Initialize()

	if err := validateNewMappingParameters(tasks); err != nil {
		panic(err)
	}
	j := jsiiProxy_Mapping{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.Mapping",
		[]interface{}{tasks},
		&j,
	)

	return &j
}

// Experimental.
func NewMapping_Override(m Mapping, tasks *[]ITask) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.Mapping",
		[]interface{}{tasks},
		m,
	)
}

// Specifies an addition mapping of two numeric values from asource to a destination.
//
// Returns: an IMapping instance.
// Experimental.
func Mapping_Add(sourceField1 *Field, sourceField2 *Field, to *Field) Mapping {
	_init_.Initialize()

	if err := validateMapping_AddParameters(sourceField1, sourceField2, to); err != nil {
		panic(err)
	}
	var returns Mapping

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.Mapping",
		"add",
		[]interface{}{sourceField1, sourceField2, to},
		&returns,
	)

	return returns
}

// A mapping definition building concatenation of source fields into a destination field.
//
// Returns: a mapping instance with concatenation definition.
// Experimental.
func Mapping_Concat(from *[]*Field, to *Field, format *string) Mapping {
	_init_.Initialize()

	if err := validateMapping_ConcatParameters(from, to, format); err != nil {
		panic(err)
	}
	var returns Mapping

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.Mapping",
		"concat",
		[]interface{}{from, to, format},
		&returns,
	)

	return returns
}

// Specifies a division mapping of two numeric values from a source to a destination.
//
// Returns: an IMapping instance.
// Experimental.
func Mapping_Divide(sourceField1 *Field, sourceField2 *Field, to *Field) Mapping {
	_init_.Initialize()

	if err := validateMapping_DivideParameters(sourceField1, sourceField2, to); err != nil {
		panic(err)
	}
	var returns Mapping

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.Mapping",
		"divide",
		[]interface{}{sourceField1, sourceField2, to},
		&returns,
	)

	return returns
}

// Experimental.
func Mapping_Map(from *Field, to *Field) IMapping {
	_init_.Initialize()

	if err := validateMapping_MapParameters(from, to); err != nil {
		panic(err)
	}
	var returns IMapping

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.Mapping",
		"map",
		[]interface{}{from, to},
		&returns,
	)

	return returns
}

// Experimental.
func Mapping_MapAll(config *MapAllConfig) IMapping {
	_init_.Initialize()

	if err := validateMapping_MapAllParameters(config); err != nil {
		panic(err)
	}
	var returns IMapping

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.Mapping",
		"mapAll",
		[]interface{}{config},
		&returns,
	)

	return returns
}

// Specifies a multiplication mapping of two numeric values from a source to a destination.
//
// Returns: an IMapping instance.
// Experimental.
func Mapping_Multiply(sourceField1 *Field, sourceField2 *Field, to *Field) Mapping {
	_init_.Initialize()

	if err := validateMapping_MultiplyParameters(sourceField1, sourceField2, to); err != nil {
		panic(err)
	}
	var returns Mapping

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.Mapping",
		"multiply",
		[]interface{}{sourceField1, sourceField2, to},
		&returns,
	)

	return returns
}

// Specifies a subtraction mapping of two numeric values from a source to a destination.
//
// Returns: an IMapping instance.
// Experimental.
func Mapping_Subtract(sourceField1 *Field, sourceField2 *Field, to *Field) Mapping {
	_init_.Initialize()

	if err := validateMapping_SubtractParameters(sourceField1, sourceField2, to); err != nil {
		panic(err)
	}
	var returns Mapping

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.Mapping",
		"subtract",
		[]interface{}{sourceField1, sourceField2, to},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mapping) Bind(flow IFlow, source ISource) *[]*awsappflow.CfnFlow_TaskProperty {
	if err := m.validateBindParameters(flow, source); err != nil {
		panic(err)
	}
	var returns *[]*awsappflow.CfnFlow_TaskProperty

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{flow, source},
		&returns,
	)

	return returns
}

