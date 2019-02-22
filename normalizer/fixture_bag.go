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

package normalizer

import (
	"fmt"

	"github.com/guiyomh/charlatan/contract"
)

type FixtureBag map[string]contract.Fixture

// Add the given fixture in the map
// If a fixture of that id already existed, it will be overridden
func (fb FixtureBag) Add(f contract.Fixture) contract.Bager {
	if fb == nil {
		return FixtureBag{f.Id(): f}
	}
	fb[f.Id()] = f
	return fb
}

// Without Creates a new instance which will not contain the fixture of the given ID.
// Will still proceed even if such fixtures does not exist.
func (fb FixtureBag) Without(fixture contract.Fixture) contract.Bager {
	clone := make(FixtureBag, len(fb)-1)
	for k, v := range fb {
		if k == fixture.Id() {
			continue
		}
		clone[k] = v
	}
	return clone
}

// MergeWith creates a new instance with values of the two FixtureBag
func (fb FixtureBag) MergeWith(bag contract.Bager) (contract.Bager, error) {
	clone := make(FixtureBag)
	for k, v := range fb {
		clone[k] = v
	}
	newBag, ok := bag.(FixtureBag)
	if !ok {
		return clone, fmt.Errorf("Could not convert %T to a FixtureBag", bag)
	}

	for k, v := range newBag {
		clone[k] = v
	}
	return clone, nil
}

// Has check if a fixture with the id exist in the bag
func (fb FixtureBag) Has(id string) bool {
	_, ok := fb[id]
	return ok
}

// Get retrieve the fixture that matching the id
func (fb FixtureBag) Get(id string) (contract.Fixture, error) {
	if fb.Has(id) {
		return fb[id], nil
	}
	return nil, fmt.Errorf("Could not find the fixture '%s'.", id)
}
