package cdklabscdkappflow


// Represents the OAuth flow enabled for the GA4.
// Experimental.
type GoogleAnalytics4OAuthFlow struct {
	// The details required for executing the refresh token grant flow.
	// Experimental.
	RefreshTokenGrant *GoogleAnalytics4RefreshTokenGrantFlow `field:"required" json:"refreshTokenGrant" yaml:"refreshTokenGrant"`
}

