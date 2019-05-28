package models

import "github.com/guiyomh/charlatan/contract"

type Fixture struct {
	id     string
	fields map[string]contract.Value
}

func NewFixture(id string, fields map[string]contract.Value) contract.Fixture {
	return &Fixture{
		id:     id,
		fields: fields,
	}
}

func (f *Fixture) Id() string {
	return f.id
}

func (f *Fixture) Fields() map[string]contract.Value {
	return f.fields
}
