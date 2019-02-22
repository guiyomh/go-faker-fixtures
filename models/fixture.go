package models

import "github.com/guiyomh/charlatan/contract"

type Fixture struct {
	id string
}

func NewFixture(id string) contract.Fixture {
	return &Fixture{
		id: id,
	}
}

func (f *Fixture) Id() string {
	return f.id
}
