package generator

import (
	"regexp"
	"strings"

	"github.com/guiyomh/charlatan/internal/app/model"
	internalcontracts "github.com/guiyomh/charlatan/internal/contracts"
	mids "github.com/guiyomh/charlatan/internal/pkg/generator/middleware"
	"github.com/guiyomh/charlatan/pkg/faker"
	fakercontracts "github.com/guiyomh/charlatan/pkg/faker/contracts"
	"github.com/guiyomh/charlatan/pkg/ranger"
)

var (
	objectSetRegex, _ = regexp.Compile(`(?i)^(?P<record>[a-z0-9-_]+)(?P<quantifier>\{[a-z0-9\.,]+\})?( \(((?P<isTemplate>template)|(extends (?P<template>[a-z0-9-_]+))\)))?`)
	myRanger          = ranger.NewRanger()
)

type Generator struct {
	faker fakercontracts.Faker
}

// NewGenerator factory to create a Generator
func NewGenerator() *Generator {
	return &Generator{
		faker: faker.NewValue(),
	}
}

// GenerateRecords build records from fixture
func (g Generator) GenerateRecords(data model.FixtureTables) ([]internalcontracts.Row, error) {
	tpls, recordSets := g.classifyData(data)
	rows, err := g.buildRecord(tpls, recordSets)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (g Generator) classifyData(tbls model.FixtureTables) (map[string]*model.Template, []internalcontracts.RowSet) {
	tpls := make(map[string]*model.Template)
	objs := make([]internalcontracts.RowSet, 0)
	for tableName, records := range tbls {
		for recordName, fields := range records {

			groups := objectSetRegex.FindAllStringSubmatch(recordName, -1)[0]
			name := groups[1]
			rangeRef := groups[2]
			isTemplate := groups[4] == "template"
			hasExtend := groups[7] != ""
			parent := groups[7]
			if isTemplate {
				tpls[name] = model.NewTemplate(name, fields)
			} else {
				objectSet := model.NewObjectSet(tableName, name, fields, hasExtend, rangeRef, parent)
				if rangeRef != "" {
					for _, set := range myRanger.BuildRecordName(name, rangeRef) {
						objectSet.AddRangeRowReference(set)
					}
				} else {
					objectSet.AddRangeRowReference(objectSet.Name())

				}
				objs = append(objs, objectSet)
			}
		}
	}
	return tpls, objs
}

func (g Generator) buildRecord(templates map[string]*model.Template, recordSets []internalcontracts.RowSet) ([]internalcontracts.Row, error) {
	rows := make([]internalcontracts.Row, 0)
	for _, objectSet := range recordSets {

		if objectSet.HasExtend() {
			objectSet = g.completeField(templates, objectSet)
		}
		rowSets := g.createRows(objectSet)
		rows = append(rows, rowSets...)
	}
	return rows, nil
}

func (g Generator) completeField(templates map[string]*model.Template, set internalcontracts.RowSet) internalcontracts.RowSet {
	for fieldName, value := range templates[set.ParentName()].Fields {
		set.AddField(fieldName, value)
	}
	return set
}

func (g Generator) createRows(objectSet internalcontracts.RowSet) []internalcontracts.Row {
	rows := make([]internalcontracts.Row, 0)
	for _, rowReference := range objectSet.RangeRowReference() {
		row := g.createRow(rowReference, objectSet)
		rows = append(rows, row)
	}
	return rows
}

func (g Generator) createRow(rowReference string, objectSet internalcontracts.RowSet) internalcontracts.Row {

	current := strings.Replace(rowReference, objectSet.Name(), "", 1)
	row := model.NewRow(rowReference, objectSet.TableName())
	for field, value := range objectSet.Fields() {
		v := g.generateValue(current, value)
		row.AddField(field, v)
		g.getDependency(field, v, row)
	}
	return row
}

func (g Generator) applyMiddleWare(value interface{}, mids ...middleware) interface{} {
	for _, m := range mids {
		value = m(value)
	}
	return value
}

func (g Generator) getDependency(field string, value interface{}, row internalcontracts.Row) {
	if _, ok := value.(string); !ok {
		return
	}
	relation, err := model.NewRelation(value.(string))
	if err != nil {
		return
	}
	row.AddDependency(field, relation)
}

func (g Generator) generateValue(current string, value interface{}) interface{} {

	chain := make([]middleware, 0)
	chain = append(chain, mids.CurrentMiddleware(current))
	chain = append(chain, mids.FakerMiddleware(g.faker))

	value = g.applyMiddleWare(value, chain...)

	return value
}
