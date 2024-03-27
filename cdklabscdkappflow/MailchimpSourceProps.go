package cdklabscdkappflow


// Experimental.
type MailchimpSourceProps struct {
	// Experimental.
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Experimental.
	Profile MailchimpConnectorProfile `field:"required" json:"profile" yaml:"profile"`
}

