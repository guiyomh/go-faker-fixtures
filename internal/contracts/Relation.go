package contracts

// Relation represents a relation to an another record
type Relation interface {
	RecordName() string
	FieldName() string
}
