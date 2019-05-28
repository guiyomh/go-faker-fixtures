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
	"github.com/sarulabs/di"
)

var Services []di.Def = []di.Def{
	{
		Name:  "app.normalizer.list",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &List{}, nil
		},
	},
	{
		Name:  "app.normalizer.range",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &Range{}, nil
		},
	},
	{
		Name:  "app.normalizer.registry",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			normalizers := make([]contract.Chainabler, 0)
			if n, ok := ctn.Get("app.normalizer.range").(contract.Chainabler); ok {
				normalizers = append(normalizers, n)
			}
			if n, ok := ctn.Get("app.normalizer.list").(contract.Chainabler); ok {
				normalizers = append(normalizers, n)
			}
			if len(normalizers) == 0 {
				return nil, fmt.Errorf("The service app.normalizer.registry needs Chainabler to works")
			}
			registry := NewRegistry(normalizers)
			return registry, nil
		},
	},
}
