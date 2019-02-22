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

package loader

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

type Locatorer interface {
	LocateFiles(dir string) (files []string, err error)
}

type FixtureLocator struct {
	fs afero.Fs
}

func NewFixtureLocator(fs afero.Fs) *FixtureLocator {
	if fs == nil {
		fs = afero.NewOsFs()
	}
	return &FixtureLocator{
		fs: fs,
	}
}

// LocateFiles all the fixture files to load
func (fl *FixtureLocator) LocateFiles(root string) ([]string, error) {
	log.WithFields(log.Fields{"root": root}).Info("Locate fixtures file")
	defer func() {
		log.Info("Fixtures file located")
	}()
	var files []string

	err := afero.Walk(fl.fs, root, func(dir string, info os.FileInfo, err error) error {
		if !isValidFixture(info) {
			return nil
		}
		files = append(files, dir)
		return nil
	})
	if err != nil {
		return files, err
	}

	return files, nil
}

func isValidFixture(f os.FileInfo) bool {
	return !f.IsDir() && (strings.HasSuffix(f.Name(), ".yml") || strings.HasSuffix(f.Name(), ".yaml") && f.Size() > 0)
}
