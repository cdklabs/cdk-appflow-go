package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type MicrosoftDynamics365OAuthSettings struct {
	// The access token to be used when interacting with Microsoft Dynamics 365.
	//
	// Note that if only the access token is provided AppFlow is not able to retrieve a fresh access token when the current one is expired.
	// Experimental.
	AccessToken awscdk.SecretValue `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Experimental.
	Endpoints *MicrosoftDynamics365OAuthEndpointsSettings `field:"optional" json:"endpoints" yaml:"endpoints"`
	// Experimental.
	Flow *MicrosoftDynamics365OAuthFlow `field:"optional" json:"flow" yaml:"flow"`
}

