package cdklabscdkappflow


// Experimental.
type MicrosoftSharepointOnlineOAuthSettings struct {
	// Experimental.
	Endpoints *MicrosoftSharepointOnlineOAuthEndpointsSettings `field:"required" json:"endpoints" yaml:"endpoints"`
	// The access token to be used when interacting with Google Analytics 4.
	//
	// Note that if only the access token is provided AppFlow is not able to retrieve a fresh access token when the current one is expired.
	// Experimental.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Experimental.
	Flow *MicrosoftSharepointOnlineOAuthFlow `field:"optional" json:"flow" yaml:"flow"`
}

