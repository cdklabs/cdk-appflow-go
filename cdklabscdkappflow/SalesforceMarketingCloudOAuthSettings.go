package cdklabscdkappflow


// Experimental.
type SalesforceMarketingCloudOAuthSettings struct {
	// Experimental.
	Endpoints *SalesforceMarketingCloudOAuthEndpoints `field:"required" json:"endpoints" yaml:"endpoints"`
	// Experimental.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Experimental.
	Flow *SalesforceMarketingCloudFlowSettings `field:"optional" json:"flow" yaml:"flow"`
}

