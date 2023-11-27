package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type MicrosoftSharepointOnlineOAuthSettings struct {
	// The access token to be used when interacting with Microsoft Sharepoint Online.
	//
	// Note that if only the access token is provided AppFlow is not able to retrieve a fresh access token when the current one is expired.
	// Experimental.
	AccessToken awscdk.SecretValue `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Experimental.
	Endpoints *MicrosoftSharepointOnlineOAuthEndpointsSettings `field:"optional" json:"endpoints" yaml:"endpoints"`
	// Experimental.
	Flow *MicrosoftSharepointOnlineOAuthFlow `field:"optional" json:"flow" yaml:"flow"`
}

