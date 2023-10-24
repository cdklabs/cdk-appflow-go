package cdklabscdkappflow


// Snowflake authorization settings required for the profile.
// Experimental.
type SnowflakeBasicAuthSettings struct {
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

