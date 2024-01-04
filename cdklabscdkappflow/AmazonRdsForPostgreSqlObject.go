package cdklabscdkappflow


// The definition of the Amazon AppFlow object for Amazon RDS for PostgreSQL.
// Experimental.
type AmazonRdsForPostgreSqlObject struct {
	// PostgreSQL schema name of the table.
	// Experimental.
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// PostgreSQL table name.
	// Experimental.
	Table *string `field:"required" json:"table" yaml:"table"`
}

