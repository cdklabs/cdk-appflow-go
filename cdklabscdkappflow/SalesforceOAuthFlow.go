package cdklabscdkappflow


// Experimental.
type SalesforceOAuthFlow struct {
	// The parameters required for the refresh token grant OAuth flow.
	// Experimental.
	RefreshTokenGrant *SalesforceOAuthRefreshTokenGrantFlow `field:"optional" json:"refreshTokenGrant" yaml:"refreshTokenGrant"`
	// The parameters required for the refresh token grant OAuth flow.
	// Deprecated: - this property will be removed in the future releases. Use refreshTokenGrant property instead.
	RefresTokenGrant *SalesforceOAuthRefreshTokenGrantFlow `field:"optional" json:"refresTokenGrant" yaml:"refresTokenGrant"`
}

