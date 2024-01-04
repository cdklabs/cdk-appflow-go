package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Basic authentication settings for the JdbcSmallDataScaleConnectorProfile.
// Experimental.
type JdbcSmallDataScaleBasicAuthSettings struct {
	// The password of the identity used for interacting with the database.
	// Experimental.
	Password awscdk.SecretValue `field:"required" json:"password" yaml:"password"`
	// The username of the identity used for interacting with the database.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

