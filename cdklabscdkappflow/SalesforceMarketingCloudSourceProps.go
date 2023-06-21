package cdklabscdkappflow


// Properties of a Salesforce Marketing Cloud Source.
// Experimental.
type SalesforceMarketingCloudSourceProps struct {
	// Experimental.
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Experimental.
	Profile SalesforceMarketingCloudConnectorProfile `field:"required" json:"profile" yaml:"profile"`
}

