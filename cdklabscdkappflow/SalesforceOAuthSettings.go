package cdklabscdkappflow


// Experimental.
type SalesforceOAuthSettings struct {
	// Experimental.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Experimental.
	Flow *SalesforceOAuthFlow `field:"optional" json:"flow" yaml:"flow"`
}

