package cdklabscdkappflow


// The OAuth elements required for the execution of the refresh token grant flow.
// Experimental.
type GoogleAnalytics4RefreshTokenGrantFlow struct {
	// The id of the client app.
	// Experimental.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The secret of the client app.
	// Experimental.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// A non-expired refresh token.
	// Experimental.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

