package cdklabscdkappflow


// Properties of a Hubspot Source.
// Experimental.
type HubSpotSourceProps struct {
	// Experimental.
	ApiVersion HubSpotApiVersion `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Experimental.
	Entity *[]*string `field:"required" json:"entity" yaml:"entity"`
	// Experimental.
	Profile HubSpotConnectorProfile `field:"required" json:"profile" yaml:"profile"`
}

