package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Experimental.
type SAPOdataConnectorProfileProps struct {
	// TODO: think if this should be here as not all connector profiles have that.
	// Experimental.
	Key awskms.IKey `field:"optional" json:"key" yaml:"key"`
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Experimental.
	ApplicationHostUrl *string `field:"required" json:"applicationHostUrl" yaml:"applicationHostUrl"`
	// Experimental.
	ApplicationServicePath *string `field:"required" json:"applicationServicePath" yaml:"applicationServicePath"`
	// Experimental.
	ClientNumber *string `field:"required" json:"clientNumber" yaml:"clientNumber"`
	// Experimental.
	LogonLanguage *string `field:"required" json:"logonLanguage" yaml:"logonLanguage"`
	// Experimental.
	BasicAuth *SAPOdataBasicAuthSettings `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	// Experimental.
	OAuth *SAPOdataOAuthSettings `field:"optional" json:"oAuth" yaml:"oAuth"`
	// Experimental.
	PortNumber *float64 `field:"optional" json:"portNumber" yaml:"portNumber"`
}

