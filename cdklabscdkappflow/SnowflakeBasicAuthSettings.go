package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Snowflake authorization settings required for the profile.
// Experimental.
type SnowflakeBasicAuthSettings struct {
	// Experimental.
	Password awscdk.SecretValue `field:"required" json:"password" yaml:"password"`
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

