package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type ServiceNowBasicSettings struct {
	// Experimental.
	Password awscdk.SecretValue `field:"required" json:"password" yaml:"password"`
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

