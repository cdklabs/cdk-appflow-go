package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type HubSpotOAuthSettings struct {
	// The access token to be used when interacting with Hubspot.
	//
	// Note that if only the access token is provided AppFlow is not able to retrieve a fresh access token when the current one is expired.
	// Default: Retrieves a fresh accessToken with the information in the [flow property]{@link HubSpotOAuthSettings#flow }.
	//
	// Experimental.
	AccessToken awscdk.SecretValue `field:"optional" json:"accessToken" yaml:"accessToken"`
	// The OAuth token and authorization endpoints.
	// Experimental.
	Endpoints *HubSpotOAuthEndpoints `field:"optional" json:"endpoints" yaml:"endpoints"`
	// The OAuth flow used for obtaining a new accessToken when the old is not present or expired.
	// Default: undefined. AppFlow will not request any new accessToken after expiry.
	//
	// Experimental.
	Flow *HubSpotOAuthFlow `field:"optional" json:"flow" yaml:"flow"`
}

