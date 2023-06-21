package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Experimental.
type ServiceNowConnectorProfileProps struct {
	// TODO: think if this should be here as not all connector profiles have that.
	// Experimental.
	Key awskms.IKey `field:"optional" json:"key" yaml:"key"`
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Experimental.
	BasicAuth *ServiceNowBasicSettings `field:"required" json:"basicAuth" yaml:"basicAuth"`
	// Experimental.
	InstanceUrl *string `field:"required" json:"instanceUrl" yaml:"instanceUrl"`
}

