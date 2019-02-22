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

package fixture

import "github.com/guiyomh/charlatan/contract"

type FixtureSet struct {
	bag contract.Bager
}

func (f *FixtureSet) Fixtures() contract.Bager {
	return f.bag
}

type fixtureSet struct {
	Name   string
	Fields map[string]field
}

func newFixtureSets(data map[string]map[string]interface{}) map[string]fixtureSet {

	sets := make(map[string]fixtureSet, len(data))
	for k, m := range data {
		sets[k] = newFixtureSet(k, m)
	}
	return sets
}

func newFixtureSet(name string, data map[string]interface{}) fixtureSet {
	fields := make(map[string]field, len(data))
	var f field
	for k, v := range data {
		f = field{
			Name:  k,
			Value: v,
		}
		fields[k] = f
	}
	return fixtureSet{
		Name:   name,
		Fields: fields,
	}
}
