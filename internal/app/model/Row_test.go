package model

import (
	"testing"

	internalcontracts "github.com/guiyomh/charlatan/internal/contracts"
	"github.com/stretchr/testify/assert"
)

func TestNewRow(t *testing.T) {
	want := &row{
		name:                "foo",
		tableName:           "bar",
		fields:              make(map[string]interface{}),
		dependencyReference: make(map[string]internalcontracts.Relation, 0),
		pk:                  nil,
	}

	got := NewRow("foo", "bar")
	assert.Equal(t, want, got)
}

func mustRelation(relation internalcontracts.Relation, err error) internalcontracts.Relation {
	if err != nil {
		return nil
	}
	return relation
}

func TestRow_HasDependencyOf(t *testing.T) {
	type fields struct {
		Name                string
		TableName           string
		DependencyReference map[string]internalcontracts.Relation
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"Testing with zero relation",
			fields{
				Name:                "foo",
				TableName:           "bar",
				DependencyReference: map[string]internalcontracts.Relation{},
			},
			args{name: "foo"},
			false,
		},
		{
			"Testing one dependency with two relations",
			fields{
				Name:      "foo",
				TableName: "bar",
				DependencyReference: map[string]internalcontracts.Relation{
					"r1": mustRelation(NewRelation("@foo.bar")),
					"r2": mustRelation(NewRelation("@bar.foo")),
				},
			},
			args{name: "foo"},
			true,
		},
		{
			"Testing zero dependency with two relations",
			fields{
				Name:      "foo",
				TableName: "bar",
				DependencyReference: map[string]internalcontracts.Relation{
					"r1": mustRelation(NewRelation("@foo.bar")),
					"r2": mustRelation(NewRelation("@bar.foo")),
				},
			},
			args{name: "harry"},
			false,
		},
	}
	for _, tt := range tests {
		r := &row{
			name:                tt.fields.Name,
			tableName:           tt.fields.TableName,
			dependencyReference: tt.fields.DependencyReference,
		}
		assert.Equal(t, tt.want, r.hasDependencyOf(tt.args.name))
	}
}

func TestRow_HasDependencies(t *testing.T) {
	type fields struct {
		Name                string
		TableName           string
		DependencyReference map[string]internalcontracts.Relation
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"Testing with zero relation",
			fields{
				Name:                "foo",
				TableName:           "bar",
				DependencyReference: map[string]internalcontracts.Relation{},
			},
			false,
		},
		{
			"Testing one dependency with two relations",
			fields{
				Name:      "foo",
				TableName: "bar",
				DependencyReference: map[string]internalcontracts.Relation{
					"r1": mustRelation(NewRelation("foo.bar")),
					"r2": mustRelation(NewRelation("bar.foo")),
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		r := &row{
			name:                tt.fields.Name,
			tableName:           tt.fields.TableName,
			dependencyReference: tt.fields.DependencyReference,
		}
		assert.Equal(t, tt.want, r.HasDependencies())
	}
}
