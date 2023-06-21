package cdklabscdkappflow


// Experimental.
type S3OutputFormatting struct {
	// Experimental.
	Aggregation *S3FileAggregation `field:"optional" json:"aggregation" yaml:"aggregation"`
	// Experimental.
	FilePrefix *S3OutputFilePrefix `field:"optional" json:"filePrefix" yaml:"filePrefix"`
	// Experimental.
	FileType S3OutputFileType `field:"optional" json:"fileType" yaml:"fileType"`
	// Experimental.
	PreserveSourceDataTypes *bool `field:"optional" json:"preserveSourceDataTypes" yaml:"preserveSourceDataTypes"`
}

