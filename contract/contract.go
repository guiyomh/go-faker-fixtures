// Copyright (C) 2019 Guillaume CAMUS
//
// This file is part of Charlatan.
//
// Charlatan is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Charlatan is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Charlatan.  If not, see <http://www.gnu.org/licenses/>.

package contract

type Identifier interface {
	Id() string
}

type Value interface {
	String() string
	Value() interface{}
}

// Fixture is a value struct representing an raw to be built.
type Fixture interface {
	Identifier
	Fields() map[string]Value
}

// Denormalizer denormalizes the parsed data into a comprehensive collection of fixtures.
type Normalizer interface {
	Denormalize(bag Bager, ref string, data map[string]interface{}) (Bager, error)
}

type Register interface {
	Denormalize(table string, data map[string]map[string]interface{}) (Bager, error)
}

type Chainabler interface {
	Normalizer
	CanDenormalize(ref string) bool
}

type Collectioner interface {
	Chainabler
	BuildIds(matches map[string]string) ([]string, error)
}

type Bager interface {
	Add(f Fixture) Bager
	Without(fixture Fixture) Bager
	Remove(key string) Bager
	MergeWith(newFixture Bager) (Bager, error)
	Clone() Bager
	Has(id string) bool
	Get(id string) (Fixture, error)
}
