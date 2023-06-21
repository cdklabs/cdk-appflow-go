package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"
)

// Experimental.
type WriteOperation interface {
	// Experimental.
	Ids() *[]*string
	// Experimental.
	Type() WriteOperationType
}

// The jsii proxy struct for WriteOperation
type jsiiProxy_WriteOperation struct {
	_ byte // padding
}

func (j *jsiiProxy_WriteOperation) Ids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WriteOperation) Type() WriteOperationType {
	var returns WriteOperationType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewWriteOperation(type_ WriteOperationType, ids *[]*string) WriteOperation {
	_init_.Initialize()

	if err := validateNewWriteOperationParameters(type_); err != nil {
		panic(err)
	}
	j := jsiiProxy_WriteOperation{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.WriteOperation",
		[]interface{}{type_, ids},
		&j,
	)

	return &j
}

// Experimental.
func NewWriteOperation_Override(w WriteOperation, type_ WriteOperationType, ids *[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.WriteOperation",
		[]interface{}{type_, ids},
		w,
	)
}

// Experimental.
func WriteOperation_Delete(ids *[]*string) WriteOperation {
	_init_.Initialize()

	var returns WriteOperation

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.WriteOperation",
		"delete",
		[]interface{}{ids},
		&returns,
	)

	return returns
}

// Experimental.
func WriteOperation_Insert(ids *[]*string) WriteOperation {
	_init_.Initialize()

	var returns WriteOperation

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.WriteOperation",
		"insert",
		[]interface{}{ids},
		&returns,
	)

	return returns
}

// Experimental.
func WriteOperation_Update(ids *[]*string) WriteOperation {
	_init_.Initialize()

	var returns WriteOperation

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.WriteOperation",
		"update",
		[]interface{}{ids},
		&returns,
	)

	return returns
}

// Experimental.
func WriteOperation_Upsert(ids *[]*string) WriteOperation {
	_init_.Initialize()

	var returns WriteOperation

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-appflow.WriteOperation",
		"upsert",
		[]interface{}{ids},
		&returns,
	)

	return returns
}

