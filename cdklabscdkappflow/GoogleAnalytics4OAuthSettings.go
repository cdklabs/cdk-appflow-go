package cdklabscdkappflow


// Experimental.
type GoogleAnalytics4OAuthSettings struct {
	// The access token to be used when interacting with Google Analytics 4.
	//
	// Note that if only the access token is provided AppFlow is not able to retrieve a fresh access token when the current one is expired.
	// Experimental.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Experimental.
	Endpoints *GoogleAnalytics4OAuthEndpointsSettings `field:"optional" json:"endpoints" yaml:"endpoints"`
	// Experimental.
	Flow *GoogleAnalytics4OAuthFlow `field:"optional" json:"flow" yaml:"flow"`
}

