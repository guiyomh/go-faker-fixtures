package model

import "github.com/guiyomh/charlatan/internal/contracts"

// ObjectSet contains metadatas necessary to build row
type objectSet struct {
	tableName           string
	hasExtend           bool
	name                string
	rangeReference      string
	parentName          string
	fields              map[string]interface{}
	rangeRowReference   []string
	dependencyReference []string
}

// NewObjectSet is a factory to create an objectSet
func NewObjectSet(
	tableName string,
	name string,
	fields map[string]interface{},
	hasExtend bool,
	rangeReference string,
	parentName string,
) contracts.RowSet {
	return &objectSet{
		tableName:           tableName,
		hasExtend:           hasExtend,
		fields:              fields,
		name:                name,
		rangeReference:      rangeReference,
		parentName:          parentName,
		rangeRowReference:   make([]string, 0),
		dependencyReference: make([]string, 0),
	}
}

func (o *objectSet) TableName() string {
	return o.tableName
}

func (o *objectSet) HasExtend() bool {
	return o.hasExtend
}

func (o *objectSet) Name() string {
	return o.name
}

func (o *objectSet) ParentName() string {
	return o.parentName
}

func (o *objectSet) RangeReference() string {
	return o.rangeReference
}

func (o *objectSet) Fields() map[string]interface{} {
	return o.fields
}

func (o *objectSet) AddField(fieldName string, value interface{}) {
	o.fields[fieldName] = value
}

func (o *objectSet) RangeRowReference() []string {
	return o.rangeRowReference
}

func (o *objectSet) AddRangeRowReference(rangeRef string) {
	o.rangeRowReference = append(o.rangeRowReference, rangeRef)
}

func (o *objectSet) DependencyReference() []string {
	return o.dependencyReference
}

func (o *objectSet) AddDependencyReference(reference string) {
	o.dependencyReference = append(o.dependencyReference, reference)
}
