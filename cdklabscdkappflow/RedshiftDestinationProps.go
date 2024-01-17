package cdklabscdkappflow


// Experimental.
type RedshiftDestinationProps struct {
	// A Redshift table object (optionally with the schema).
	// Experimental.
	Object *RedshiftDestinationObject `field:"required" json:"object" yaml:"object"`
	// An instance of the.
	// Experimental.
	Profile RedshiftConnectorProfile `field:"required" json:"profile" yaml:"profile"`
	// The settings that determine how Amazon AppFlow handles an error when placing data in the Salesforce destination.
	//
	// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure.
	// Experimental.
	ErrorHandling *ErrorHandlingConfiguration `field:"optional" json:"errorHandling" yaml:"errorHandling"`
}

