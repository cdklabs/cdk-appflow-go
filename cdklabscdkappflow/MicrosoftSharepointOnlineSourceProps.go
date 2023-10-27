package cdklabscdkappflow


// Properties of a Microsoft Sharepoint Online Source.
// Experimental.
type MicrosoftSharepointOnlineSourceProps struct {
	// Experimental.
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Experimental.
	Object *MicrosoftSharepointOnlineObject `field:"required" json:"object" yaml:"object"`
	// Experimental.
	Profile MicrosoftSharepointOnlineConnectorProfile `field:"required" json:"profile" yaml:"profile"`
}

