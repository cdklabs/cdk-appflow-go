package cdklabscdkappflow


// Properties of a Google Analytics v4 Source.
// Experimental.
type MicrosoftSharepointOnlineSourceProps struct {
	// Experimental.
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Experimental.
	Profile MicrosoftSharepointOnlineConnectorProfile `field:"required" json:"profile" yaml:"profile"`
}

