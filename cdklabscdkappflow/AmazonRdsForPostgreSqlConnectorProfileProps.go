package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties of the AmazonRdsForPostgreSqlConnectorProfile.
// Experimental.
type AmazonRdsForPostgreSqlConnectorProfileProps struct {
	// TODO: think if this should be here as not all connector profiles have that.
	// Experimental.
	Key awskms.IKey `field:"optional" json:"key" yaml:"key"`
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The auth settings for the profile.
	// Experimental.
	BasicAuth *AmazonRdsForPostgreSqlBasicAuthSettings `field:"required" json:"basicAuth" yaml:"basicAuth"`
	// The name of the PostgreSQL database.
	// Experimental.
	Database *string `field:"required" json:"database" yaml:"database"`
	// The PostgreSQL hostname.
	// Experimental.
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// The PostgreSQL communication port.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

