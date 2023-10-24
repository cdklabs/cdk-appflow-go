package cdklabscdkappflow


// Properties that are required to create a Snowflake destination.
// Experimental.
type SnowflakeDestinationProps struct {
	// A Snowflake table object (optionally with the schema).
	// Experimental.
	Object *SnowflakeDestinationObject `field:"required" json:"object" yaml:"object"`
	// A Snowflake connector profile instance.
	// Experimental.
	Profile SnowflakeConnectorProfile `field:"required" json:"profile" yaml:"profile"`
	// The settings that determine how Amazon AppFlow handles an error when placing data in the Salesforce destination.
	//
	// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure.
	// Experimental.
	ErrorHandling *ErrorHandlingConfiguration `field:"optional" json:"errorHandling" yaml:"errorHandling"`
}

