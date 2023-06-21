package cdklabscdkappflow


// Experimental.
type ZendeskOAuthSettings struct {
	// Experimental.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Experimental.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Experimental.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
}

