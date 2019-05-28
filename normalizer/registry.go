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
	bag := CreateBag(data)
	log.WithFields(log.Fields{
		"registry": true,
		"table":    table,
	}).Debug("Registry start denormalize")

	var hasFoundDenorm bool
	for setname, fields := range data {
		hasFoundDenorm = false
		log.WithField("set", setname).Debugf("Searching normalizer for '%s'", setname)
		for _, denorm := range r.normalizers {
			ok := denorm.CanDenormalize(setname)
			if ok {
				hasFoundDenorm = true
				log.WithFields(log.Fields{
					"setname":    setname,
					"normalizer": fmt.Sprintf("%T", denorm),
				}).Debugf("Apply the normalizer '%T' to '%s'", denorm, setname)
				bag, err := denorm.Denormalize(bag, setname, fields)
				if err != nil {
					return bag, err
				}
				break
			}
		}
		if hasFoundDenorm == false {
			log.WithField("set", setname).Warnf("No normalizer was found for '%s'", setname)
		}
	}
	log.WithFields(log.Fields{
		"registry": true,
		"table":    table,
	}).Debug("Registry end denormalize")
	return bag, nil
}
