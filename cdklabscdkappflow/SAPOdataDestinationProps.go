package cdklabscdkappflow


// Experimental.
type SAPOdataDestinationProps struct {
	// The SAPOdata object for which the operation is to be set.
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Experimental.
	Operation WriteOperation `field:"required" json:"operation" yaml:"operation"`
	// Experimental.
	Profile SAPOdataConnectorProfile `field:"required" json:"profile" yaml:"profile"`
	// The settings that determine how Amazon AppFlow handles an error when placing data in the SAPOdata destination.
	//
	// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure.
	// Experimental.
	ErrorHandling *ErrorHandlingConfiguration `field:"optional" json:"errorHandling" yaml:"errorHandling"`
	// Experimental.
	SuccessResponseHandling *SAPOdataSuccessResponseHandlingConfiguration `field:"optional" json:"successResponseHandling" yaml:"successResponseHandling"`
}

