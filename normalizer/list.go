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
	"regexp"
	"strings"

	"github.com/guiyomh/charlatan/contract"
	"github.com/guiyomh/charlatan/models"
)

var LIST_REGEX = regexp.MustCompile(`(?i)(?P<refbase>[\p{L}\._\/]+)\{(?P<list>[\p{L}\._\/]+(?:,\s?[^,\s]+)*)\}`)

type List struct {
}

func (l *List) CanDenormalize(ref string) bool {
	return LIST_REGEX.MatchString(ref)
}

// BuildIds parse the referense to build a range of ids
// For example user_{bob,alice} => []string{
// 	"user_bob",
// 	"user_alice",
// }
func (l *List) BuildIds(matches map[string]string) ([]string, error) {
	var ids = make([]string, 0)
	refbase, ok := matches["refbase"]
	if !ok {
		return ids, fmt.Errorf("Could not retrieve 'refbase' in '%v'", matches)
	}
	list, ok := matches["list"]
	if !ok {
		return ids, fmt.Errorf("Could not retrieve 'list' in '%v'", matches)
	}
	for _, s := range strings.Split(list, ",") {
		ids = append(ids, fmt.Sprintf("%s%s", refbase, strings.TrimSpace(s)))
	}
	return ids, nil
}

func (l *List) Denormalize(ref string, data map[string]interface{}) (contract.Bager, error) {
	bag := make(FixtureBag, 0)
	matches := LIST_REGEX.FindStringSubmatch(ref)
	if matches == nil {
		return bag, fmt.Errorf("'%v' could not match the pattern %v", ref, LIST_REGEX.String())
	}
	result := make(map[string]string, len(matches)-1)
	limit := len(matches)
	for i, name := range LIST_REGEX.SubexpNames() {
		if i != 0 && i < limit && len(name) > 0 && len(matches[i]) > 0 {
			result[name] = matches[i]
		}
	}
	ids, err := l.BuildIds(result)
	if err != nil {
		return bag, err
	}
	var f contract.Fixture
	for _, id := range ids {
		f = models.NewFixture(id)
		bag.Add(f)
	}
	return bag, nil
}
