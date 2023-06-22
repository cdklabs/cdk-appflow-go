package cdklabscdkappflow


// Experimental.
type SAPOdataOAuthSettings struct {
	// Experimental.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Experimental.
	Endpoints *SAPOdataOAuthEndpoints `field:"optional" json:"endpoints" yaml:"endpoints"`
	// Experimental.
	Flow *SAPOdataOAuthFlows `field:"optional" json:"flow" yaml:"flow"`
}

