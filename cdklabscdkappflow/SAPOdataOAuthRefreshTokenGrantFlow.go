package cdklabscdkappflow


// Experimental.
type SAPOdataOAuthRefreshTokenGrantFlow struct {
	// Experimental.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Experimental.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Experimental.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

