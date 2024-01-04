package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for the JdbcSmallDataScaleConnectorProfile.
// Experimental.
type JdbcSmallDataScaleConnectorProfileProps struct {
	// TODO: think if this should be here as not all connector profiles have that.
	// Experimental.
	Key awskms.IKey `field:"optional" json:"key" yaml:"key"`
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The auth settings for the profile.
	// Experimental.
	BasicAuth *JdbcSmallDataScaleBasicAuthSettings `field:"required" json:"basicAuth" yaml:"basicAuth"`
	// The name of the database.
	// Experimental.
	Database *string `field:"required" json:"database" yaml:"database"`
	// The driver for the database.
	//
	// Effectively specifies the type of database.
	// Experimental.
	Driver JdbcDriver `field:"required" json:"driver" yaml:"driver"`
	// The hostname of the database to interact with.
	// Experimental.
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// The database communication port.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

