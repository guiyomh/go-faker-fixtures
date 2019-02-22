package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {

	yaml := `
foo:
    fooset:
        field_1: toto
        field_2: 1
        field_3: true
`
	p := NewYamlParser()
	got, err := p.Parse([]byte(yaml))
	assert.NoError(t, err)
	expected := &Data{
		"foo": {
			"fooset": {
				"field_1": "toto",
				"field_2": 1,
				"field_3": true,
			},
		},
	}
	assert.Equal(t, expected, got)
}

func TestParseABadYaml(t *testing.T) {
	yaml := `
foo,toto,1,true
bar,tyty,2,false
`
	p := NewYamlParser()
	_, err := p.Parse([]byte(yaml))
	assert.Error(t, err)
}
