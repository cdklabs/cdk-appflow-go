package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type SalesforceOAuthSettings struct {
	// Experimental.
	AccessToken awscdk.SecretValue `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Experimental.
	Flow *SalesforceOAuthFlow `field:"optional" json:"flow" yaml:"flow"`
}

