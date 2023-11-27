package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type MarketoOAuthSettings struct {
	// Experimental.
	Flow *MarketoOAuthFlow `field:"required" json:"flow" yaml:"flow"`
	// Experimental.
	AccessToken awscdk.SecretValue `field:"optional" json:"accessToken" yaml:"accessToken"`
}

