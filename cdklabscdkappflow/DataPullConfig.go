package cdklabscdkappflow


// Experimental.
type DataPullConfig struct {
	// Experimental.
	Mode DataPullMode `field:"required" json:"mode" yaml:"mode"`
	// The name of the field to use as a timestamp for recurring incremental flows.
	//
	// The default field is set per particular @see ISource.
	// Experimental.
	TimestampField *string `field:"optional" json:"timestampField" yaml:"timestampField"`
}

