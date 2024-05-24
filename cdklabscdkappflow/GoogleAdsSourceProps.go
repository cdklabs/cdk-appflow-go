package cdklabscdkappflow


// Properties of a Google Ads Source.
// Experimental.
type GoogleAdsSourceProps struct {
	// Experimental.
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Experimental.
	Profile GoogleAdsConnectorProfile `field:"required" json:"profile" yaml:"profile"`
}

