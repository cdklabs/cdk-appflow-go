package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type MicrosoftSharepointOnlineRefreshTokenGrantFlow struct {
	// Experimental.
	ClientId awscdk.SecretValue `field:"optional" json:"clientId" yaml:"clientId"`
	// Experimental.
	ClientSecret awscdk.SecretValue `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Experimental.
	RefreshToken awscdk.SecretValue `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

