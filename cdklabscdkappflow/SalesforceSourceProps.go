package cdklabscdkappflow


// Experimental.
type SalesforceSourceProps struct {
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Experimental.
	Profile SalesforceConnectorProfile `field:"required" json:"profile" yaml:"profile"`
	// Experimental.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Experimental.
	EnableDynamicFieldUpdate *bool `field:"optional" json:"enableDynamicFieldUpdate" yaml:"enableDynamicFieldUpdate"`
	// Experimental.
	IncludeDeletedRecords *bool `field:"optional" json:"includeDeletedRecords" yaml:"includeDeletedRecords"`
}

