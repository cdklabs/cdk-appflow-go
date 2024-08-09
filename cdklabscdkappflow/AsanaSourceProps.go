package cdklabscdkappflow


// Experimental.
type AsanaSourceProps struct {
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Experimental.
	Profile AsanaConnectorProfile `field:"required" json:"profile" yaml:"profile"`
	// Experimental.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
}

