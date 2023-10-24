package cdklabscdkappflow


// The destination table in Snowflake.
//
// The table needs to reside in the databas and schema provided in the profile.
// Experimental.
type SnowflakeDestinationObject struct {
	// The name of the table object.
	// Experimental.
	Table *string `field:"required" json:"table" yaml:"table"`
}

