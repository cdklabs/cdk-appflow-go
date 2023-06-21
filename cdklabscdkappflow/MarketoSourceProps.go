package cdklabscdkappflow


// Experimental.
type MarketoSourceProps struct {
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Experimental.
	Profile MarketoConnectorProfile `field:"required" json:"profile" yaml:"profile"`
	// Experimental.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
}

