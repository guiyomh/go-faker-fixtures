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
	log "github.com/sirupsen/logrus"
)

type Registry struct {
	normalizers []contract.Chainabler
}

func NewRegistry(denorms []contract.Chainabler) *Registry {
	return &Registry{
		normalizers: denorms,
	}
}

func (r *Registry) Denormalize(table string, data map[string]map[string]interface{}) (contract.Bager, error) {
	var bag contract.Bager = make(FixtureBag, 0)
	log.WithFields(log.Fields{
		"registry": true,
		"table":    table,
	}).Debug("Registry start denormalize")
	for setname, fields := range data {
		log.WithFields(log.Fields{
			"registry": true,
			"set":      setname,
		}).Debug("Searching normalizer for the set")
		for _, denorm := range r.normalizers {
			ok := denorm.CanDenormalize(setname)
			if ok {
				log.WithFields(log.Fields{
					"registry":   true,
					"setname":    setname,
					"normalizer": fmt.Sprintf("%T", denorm),
				}).Debug("Find a normalizer")
				bagSet, err := denorm.Denormalize(setname, fields)
				bag, err = bag.MergeWith(bagSet)
				if err != nil {
					return bag, err
				}
				break
			}
			log.WithFields(log.Fields{
				"registry":   true,
				"setname":    setname,
				"normalizer": fmt.Sprintf("%T", denorm),
			}).Warn("No Found a normalizer for ", setname)
		}
	}
	log.WithFields(log.Fields{
		"registry": true,
		"table":    table,
	}).Debug("Registry end denormalize")
	return bag, nil
}
