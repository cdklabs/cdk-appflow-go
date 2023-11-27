package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type ZendeskOAuthSettings struct {
	// Experimental.
	ClientId awscdk.SecretValue `field:"required" json:"clientId" yaml:"clientId"`
	// Experimental.
	ClientSecret awscdk.SecretValue `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Experimental.
	AccessToken awscdk.SecretValue `field:"optional" json:"accessToken" yaml:"accessToken"`
}

