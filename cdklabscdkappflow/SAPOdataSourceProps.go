package cdklabscdkappflow


// Experimental.
type SAPOdataSourceProps struct {
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Experimental.
	Profile SAPOdataConnectorProfile `field:"required" json:"profile" yaml:"profile"`
}

