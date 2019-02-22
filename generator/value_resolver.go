package generator

import (
	"github.com/guiyomh/charlatan/contract"
)

type ValueResolver interface {
	Resolve(value contract.Value, fixture contract.Fixture, fixtureSet ResolvedFixtureSet)
}
