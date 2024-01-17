package cdklabscdkappflow


// Properties of a Google BigQuery Source.
// Experimental.
type GoogleBigQuerySourceProps struct {
	// Experimental.
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Experimental.
	Object *GoogleBigQueryObject `field:"required" json:"object" yaml:"object"`
	// Experimental.
	Profile GoogleBigQueryConnectorProfile `field:"required" json:"profile" yaml:"profile"`
}

