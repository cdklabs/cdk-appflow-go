package cdklabscdkappflow


// Represents the OAuth flow enabled for the GoogleAds.
// Experimental.
type GoogleAdsOAuthFlow struct {
	// The details required for executing the refresh token grant flow.
	// Experimental.
	RefreshTokenGrant *GoogleAdsRefreshTokenGrantFlow `field:"required" json:"refreshTokenGrant" yaml:"refreshTokenGrant"`
}

