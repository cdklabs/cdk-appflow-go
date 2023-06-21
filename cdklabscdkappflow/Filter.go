package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A representation of a mapping operation, that is an operation filtering records at the source.
// Experimental.
type Filter interface {
	OperationBase
	IFilter
	// Experimental.
	Condition() FilterCondition
	// Experimental.
	Bind(flow IFlow, source ISource) *[]*awsappflow.CfnFlow_TaskProperty
}

// The jsii proxy struct for Filter
type jsiiProxy_Filter struct {
	jsiiProxy_OperationBase
	jsiiProxy_IFilter
}

func (j *jsiiProxy_Filter) Condition() FilterCondition {
	var returns FilterCondition
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}


// Experimental.
func NewFilter(condition FilterCondition) Filter {
	_init_.Initialize()

	if err := validateNewFilterParameters(condition); err != nil {
		panic(err)
	}
	j := jsiiProxy_Filter{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.Filter",
		[]interface{}{condition},
		&j,
	)

	return &j
}

// Experimental.
func NewFilter_Override(f Filter, condition FilterCondition) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.Filter",
		[]interface{}{condition},
		f,
	)
}

// Builds a filter operation on source.
// See: FilterCondition instance.
//
// Experimental.
func Filter_When(condition FilterCondition) Filter {
	_init_.Initialize()

	if err := validateFilter_WhenParameters(condition); err != nil {
		panic(err)
	}
	var returns Filter

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.Filter",
		"when",
		[]interface{}{condition},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Filter) Bind(flow IFlow, source ISource) *[]*awsappflow.CfnFlow_TaskProperty {
	if err := f.validateBindParameters(flow, source); err != nil {
		panic(err)
	}
	var returns *[]*awsappflow.CfnFlow_TaskProperty

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{flow, source},
		&returns,
	)

	return returns
}

