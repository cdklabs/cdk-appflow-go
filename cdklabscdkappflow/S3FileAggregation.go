package cdklabscdkappflow


// Experimental.
type S3FileAggregation struct {
	// The maximum size, in MB, of the file containing portion of incoming data.
	// Experimental.
	FileSize *float64 `field:"optional" json:"fileSize" yaml:"fileSize"`
	// Experimental.
	Type S3OutputAggregationType `field:"optional" json:"type" yaml:"type"`
}

