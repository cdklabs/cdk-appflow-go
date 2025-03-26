package cdklabscdkappflow


// Represents the OAuth flow enabled for the GA4.
// Experimental.
type HubSpotOAuthFlow struct {
	// The details required for executing the refresh token grant flow.
	// Experimental.
	RefreshTokenGrant *HubSpotRefreshTokenGrantFlow `field:"required" json:"refreshTokenGrant" yaml:"refreshTokenGrant"`
}

