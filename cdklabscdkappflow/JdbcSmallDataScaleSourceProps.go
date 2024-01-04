package cdklabscdkappflow


// Experimental.
type JdbcSmallDataScaleSourceProps struct {
	// Experimental.
	Object *JdbcSmallDataScaleObject `field:"required" json:"object" yaml:"object"`
	// Experimental.
	Profile JdbcSmallDataScaleConnectorProfile `field:"required" json:"profile" yaml:"profile"`
	// Experimental.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
}

