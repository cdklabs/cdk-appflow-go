package cdklabscdkappflow


// Experimental.
type S3OutputFormatting struct {
	// Sets an aggregation approach per flow run.
	// Experimental.
	Aggregation *S3FileAggregation `field:"optional" json:"aggregation" yaml:"aggregation"`
	// Sets a prefix approach for files generated during a flow execution.
	// Experimental.
	FilePrefix *S3OutputFilePrefix `field:"optional" json:"filePrefix" yaml:"filePrefix"`
	// Sets the file type for the output files.
	// Default: - JSON file type.
	//
	// Experimental.
	FileType S3OutputFileType `field:"optional" json:"fileType" yaml:"fileType"`
	// Specifies whether AppFlow should attempt data type mapping from source when the destination output file type is Parquet.
	// Default: - do not preserve source data files.
	//
	// Experimental.
	PreserveSourceDataTypes *bool `field:"optional" json:"preserveSourceDataTypes" yaml:"preserveSourceDataTypes"`
}

