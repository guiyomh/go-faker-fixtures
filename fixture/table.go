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

package fixture

type Table struct {
	Table    string
	Metadata metadata
	Sets     map[string]fixtureSet
}

func NewTable(table string, sets map[string]map[string]interface{}) *Table {

	conn := "default"
	trunc := false

	return &Table{
		Table: table,
		Metadata: metadata{
			connection{
				Name:     conn,
				Truncate: trunc,
			},
		},
		Sets: newFixtureSets(sets),
	}
}

func (ft Table) Id() string {
	return ft.Table
}
