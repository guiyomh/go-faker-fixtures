package normalizer

import (
	"fmt"
	"regexp"

	"github.com/guiyomh/charlatan/contract"
	"github.com/guiyomh/charlatan/models"
)

type chainabler struct {
}

func (c chainabler) buildMatches(ref string, reg *regexp.Regexp) (map[string]string, error) {
	result := make(map[string]string, 0)
	matches := reg.FindStringSubmatch(ref)
	if matches == nil {
		return result, fmt.Errorf("'%v' could not match the pattern %v", ref, reg.String())
	}
	limit := len(matches)
	for i, name := range reg.SubexpNames() {
		if i != 0 && i < limit && len(name) > 0 && len(matches[i]) > 0 {
			result[name] = matches[i]
		}
	}
	return result, nil
}

func (c chainabler) buildFixture(ids []string, data map[string]interface{}) map[string]contract.Fixture {
	fields := make(map[string]contract.Value, len(data))
	fixtures := make(map[string]contract.Fixture, len(ids))
	for n, d := range data {
		fields[n] = models.NewData(d)
	}
	for _, id := range ids {
		fixtures[id] = models.NewFixture(id, fields)
	}
	return fixtures
}
