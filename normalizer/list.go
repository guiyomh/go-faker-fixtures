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
)

var LIST_REGEX = regexp.MustCompile(`(?i)(?P<refbase>[\p{L}\._\/]+)\{(?P<list>[\p{L}\._\/]+(?:,\s?[^,\s]+)*)\}`)

type List struct {
	chainabler
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
	matches, err := l.buildMatches(ref, LIST_REGEX)
	if err != nil {
		return bag, err
	}
	ids, err := l.BuildIds(matches)
	if err != nil {
		return bag, err
	}
	fixtures := l.buildFixture(ids, data)
	for _, f := range fixtures {
		bag.Add(f)
	}
	return bag, nil
}
