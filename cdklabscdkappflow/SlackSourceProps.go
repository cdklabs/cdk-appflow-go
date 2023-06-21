package cdklabscdkappflow


// Experimental.
type SlackSourceProps struct {
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Experimental.
	Profile SlackConnectorProfile `field:"required" json:"profile" yaml:"profile"`
	// Experimental.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
}

