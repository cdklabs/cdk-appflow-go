package cdklabscdkappflow


// Experimental.
type EventBridgeDestinationProps struct {
	// Experimental.
	PartnerBus *string `field:"required" json:"partnerBus" yaml:"partnerBus"`
	// Experimental.
	ErrorHandling *ErrorHandlingConfiguration `field:"optional" json:"errorHandling" yaml:"errorHandling"`
}

