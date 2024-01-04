package cdklabscdkappflow


// Properties of the AmazonRdsForPostgreSqlDestination.
// Experimental.
type AmazonRdsForPostgreSqlDestinationProps struct {
	// The destination object table to write to.
	// Experimental.
	Object *AmazonRdsForPostgreSqlObject `field:"required" json:"object" yaml:"object"`
	// The profile to use with the destination.
	// Experimental.
	Profile AmazonRdsForPostgreSqlConnectorProfile `field:"required" json:"profile" yaml:"profile"`
	// The Amazon AppFlow Api Version.
	// Experimental.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// The settings that determine how Amazon AppFlow handles an error when placing data in the destination.
	//
	// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure.
	// Experimental.
	ErrorHandling *ErrorHandlingConfiguration `field:"optional" json:"errorHandling" yaml:"errorHandling"`
}

