package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An interface representing a vertex, i.e. a source or a destination of an AppFlow flow.
// Experimental.
type IVertex interface {
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
}

// The jsii proxy for IVertex
type jsiiProxy_IVertex struct {
	_ byte // padding
}

func (j *jsiiProxy_IVertex) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}

