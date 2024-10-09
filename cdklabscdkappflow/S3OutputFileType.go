package cdklabscdkappflow


// Output file type supported by Amazon S3 Destination connector.
// Experimental.
type S3OutputFileType string

const (
	// CSV file type.
	// Experimental.
	S3OutputFileType_CSV S3OutputFileType = "CSV"
	// JSON file type.
	// Experimental.
	S3OutputFileType_JSON S3OutputFileType = "JSON"
	// Parquet file type.
	// Experimental.
	S3OutputFileType_PARQUET S3OutputFileType = "PARQUET"
)

