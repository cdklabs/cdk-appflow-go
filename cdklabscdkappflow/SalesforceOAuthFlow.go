package cdklabscdkappflow


// Experimental.
type SalesforceOAuthFlow struct {
	// The parameters required for the refresh token grant OAuth flow.
	// Experimental.
	RefreshTokenGrant *SalesforceOAuthRefreshTokenGrantFlow `field:"optional" json:"refreshTokenGrant" yaml:"refreshTokenGrant"`
}

