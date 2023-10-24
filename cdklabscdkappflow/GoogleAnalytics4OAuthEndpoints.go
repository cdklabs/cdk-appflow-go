package cdklabscdkappflow


// Google's OAuth token and authorization endpoints.
// Experimental.
type GoogleAnalytics4OAuthEndpoints struct {
	// The OAuth authorization endpoint URI.
	// Experimental.
	Authorization *string `field:"optional" json:"authorization" yaml:"authorization"`
	// The OAuth token endpoint URI.
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
}

