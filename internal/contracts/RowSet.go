package contracts

// RowSet contains metadatas necessary to build row
type RowSet interface {
	TableName() string
	HasExtend() bool
	Name() string
	RangeReference() string
	ParentName() string
	Fields() map[string]interface{}
	AddField(fieldName string, value interface{})
	RangeRowReference() []string
	AddRangeRowReference(rangeRef string)
	DependencyReference() []string
	AddDependencyReference(reference string)
}
