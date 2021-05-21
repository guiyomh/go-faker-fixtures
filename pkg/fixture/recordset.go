package fixture

import (
	"fmt"
	"regexp"

	"github.com/imdario/mergo"
)

var (
	RSRegex, _ = regexp.Compile(`(?i)^(?P<record>[a-z0-9-_]+)(?P<quantifier>\{[a-z0-9\.,]+\})?( \(((?P<isTemplate>template)|(extends (?P<template>[a-z0-9-_]+))\)))?`)
)

type Meta struct {
	Table      string
	RecordName string
	Quantifier string
	Parent     string
	Template   bool
}

type RecordSet struct {
	Fields map[string]interface{}
	Meta   Meta
}

func (r RecordSet) HasParent() bool {
	return r.Meta.Parent != ""
}

func (r RecordSet) IsTemplate() bool {
	return r.Meta.Template
}

func template(c []RecordSet, name string) (*RecordSet, error) {
	for _, r := range c {
		if r.IsTemplate() && r.Meta.RecordName == name {
			return &r, nil
		}
	}
	return nil, fmt.Errorf("no found the template: %s", name)
}

func GenerateRecords(data map[string]interface{}) []RecordSet {
	rss := make([]RecordSet, len(data))
	for table, rds := range data {

		records, ok := rds.(map[string]interface{})
		if !ok {
			fmt.Println("warning - not convert")
			continue
		}
		for name, fds := range records {

			g := RSRegex.FindAllStringSubmatch(name, -1)[0]
			fields, ok := fds.(map[string]interface{})
			if !ok {
				continue
			}
			rs := RecordSet{
				Fields: fields,
				Meta: Meta{
					Table:      table,
					RecordName: g[1],
					Quantifier: g[2],
					Template:   g[4] == "template",
					Parent:     g[7],
				},
			}
			rss = append(rss, rs)
		}
	}
	mergeTemplateFields(rss)
	return rss
}

func mergeTemplateFields(c []RecordSet) {
	for _, r := range c {
		if !r.HasParent() {
			continue
		}
		t, err := template(c, r.Meta.Parent)
		if err != nil {
			continue
		}
		mergo.MergeWithOverwrite(&r.Fields, t.Fields)
	}
}
