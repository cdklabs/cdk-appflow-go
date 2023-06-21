package cdklabscdkappflow


// Experimental.
type ZendeskSourceProps struct {
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Experimental.
	Profile ZendeskConnectorProfile `field:"required" json:"profile" yaml:"profile"`
	// Experimental.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
}

