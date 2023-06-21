package cdklabscdkappflow


// The file type that Amazon AppFlow gets from your Amazon S3 bucket.
// Experimental.
type S3InputFileType string

const (
	// Experimental.
	S3InputFileType_CSV S3InputFileType = "CSV"
	// Experimental.
	S3InputFileType_JSON S3InputFileType = "JSON"
)

