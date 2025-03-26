package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The OAuth elements required for the execution of the refresh token grant flow.
// Experimental.
type HubSpotRefreshTokenGrantFlow struct {
	// The id of the client app.
	// Experimental.
	ClientId awscdk.SecretValue `field:"optional" json:"clientId" yaml:"clientId"`
	// The secret of the client app.
	// Experimental.
	ClientSecret awscdk.SecretValue `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// A non-expired refresh token.
	// Experimental.
	RefreshToken awscdk.SecretValue `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

