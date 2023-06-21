package cdklabscdkappflow


// Experimental.
type SlackOAuthSettings struct {
	// Experimental.
	AccessToken *string `field:"required" json:"accessToken" yaml:"accessToken"`
	// Experimental.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Experimental.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

