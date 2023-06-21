package cdklabscdkappflow


// Experimental.
type SalesforceDestinationProps struct {
	// The Salesforce object for which the operation is to be set.
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Experimental.
	Operation WriteOperation `field:"required" json:"operation" yaml:"operation"`
	// Experimental.
	Profile SalesforceConnectorProfile `field:"required" json:"profile" yaml:"profile"`
	// Specifies which Salesforce API is used by Amazon AppFlow when your flow transfers data to Salesforce.
	// Experimental.
	DataTransferApi SalesforceDataTransferApi `field:"optional" json:"dataTransferApi" yaml:"dataTransferApi"`
	// The settings that determine how Amazon AppFlow handles an error when placing data in the Salesforce destination.
	//
	// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure.
	// Experimental.
	ErrorHandling *ErrorHandlingConfiguration `field:"optional" json:"errorHandling" yaml:"errorHandling"`
}

