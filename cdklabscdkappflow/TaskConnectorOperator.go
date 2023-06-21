package cdklabscdkappflow


// A pair that represents the (typically source) connector, and a task operation to be performed in the context of the connector.
// Experimental.
type TaskConnectorOperator struct {
	// Experimental.
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// Experimental.
	Type ConnectorType `field:"optional" json:"type" yaml:"type"`
}

