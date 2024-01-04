package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Basic authentication settings for the AmazonRdsForPostgreSqlConnectorProfile.
// Experimental.
type AmazonRdsForPostgreSqlBasicAuthSettings struct {
	// The password of the identity used for interacting with the Amazon RDS for PostgreSQL.
	// Experimental.
	Password awscdk.SecretValue `field:"required" json:"password" yaml:"password"`
	// The username of the identity used for interacting with the Amazon RDS for PostgreSQL.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

