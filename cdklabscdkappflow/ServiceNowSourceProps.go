package cdklabscdkappflow


// Experimental.
type ServiceNowSourceProps struct {
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Experimental.
	Profile ServiceNowConnectorProfile `field:"required" json:"profile" yaml:"profile"`
}

