package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type SlackOAuthSettings struct {
	// Experimental.
	AccessToken awscdk.SecretValue `field:"required" json:"accessToken" yaml:"accessToken"`
	// Experimental.
	ClientId awscdk.SecretValue `field:"optional" json:"clientId" yaml:"clientId"`
	// Experimental.
	ClientSecret awscdk.SecretValue `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

