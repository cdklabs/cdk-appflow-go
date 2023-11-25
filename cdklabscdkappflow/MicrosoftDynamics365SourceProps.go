package cdklabscdkappflow


// Properties of a Microsoft Dynamics 365 Source.
// Experimental.
type MicrosoftDynamics365SourceProps struct {
	// Experimental.
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Experimental.
	Profile MicrosoftDynamics365ConnectorProfile `field:"required" json:"profile" yaml:"profile"`
}

