package cdklabscdkappflow


// Experimental.
type MarketoOAuthSettings struct {
	// Experimental.
	Flow *MarketoOAuthFlow `field:"required" json:"flow" yaml:"flow"`
	// Experimental.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
}

