package cdklabscdkappflow


// Experimental.
type HubSpotDestinationProps struct {
	// Experimental.
	ApiVersion HubSpotApiVersion `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Experimental.
	Entity *[]*string `field:"required" json:"entity" yaml:"entity"`
	// Experimental.
	Operation WriteOperation `field:"required" json:"operation" yaml:"operation"`
	// Experimental.
	Profile HubSpotConnectorProfile `field:"required" json:"profile" yaml:"profile"`
	// The settings that determine how Amazon AppFlow handles an error when placing data in the HubSpot destination.
	//
	// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure.
	// Experimental.
	ErrorHandling *ErrorHandlingConfiguration `field:"optional" json:"errorHandling" yaml:"errorHandling"`
}

