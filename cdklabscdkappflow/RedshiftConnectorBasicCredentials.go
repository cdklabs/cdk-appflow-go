package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type RedshiftConnectorBasicCredentials struct {
	// Experimental.
	Password awscdk.SecretValue `field:"optional" json:"password" yaml:"password"`
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

