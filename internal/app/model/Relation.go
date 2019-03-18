package model

import (
	"errors"
	"regexp"

	"github.com/guiyomh/charlatan/internal/contracts"
)

var relationRegex, _ = regexp.Compile(`(?i)^\@(?P<rowname>[a-z0-9_-]+)\.?(?P<fieldname>[a-z0-9_-]+)?$`)

var (
	errNotFoundRelation = errors.New("No relation found")
)

// Relation represents a relation to an another record
type relation struct {
	recordName string
	fieldName  string
}

//NewRelation create a relation structure
func NewRelation(value string) (contracts.Relation, error) {
	deps := relationRegex.FindStringSubmatch(value)
	if len(deps) <= 0 {
		return nil, errNotFoundRelation
	}
	relation := &relation{recordName: deps[1]}
	if deps[2] != "" {
		relation.fieldName = deps[2]
	}
	return relation, nil
}

func (r *relation) RecordName() string {
	return r.recordName
}

func (r *relation) FieldName() string {
	return r.fieldName
}
