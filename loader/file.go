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
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

type FileLoader struct {
	fs afero.Fs
}

func NewFileLoader(fs afero.Fs) Loader {
	return &FileLoader{
		fs: fs,
	}
}

func (fl *FileLoader) Load(fixturesFiles []string) []byte {
	log.WithFields(log.Fields{"files": fixturesFiles}).Info("Fixtures found")
	defer func() { log.Info("Fixtures loaded") }()
	size := len(fixturesFiles)
	files := make(chan string, size)
	data := make(chan []byte, size)

	//init parser
	for w := 1; w <= 4; w++ {
		go fl.parse(w, files, data)
	}
	//start jobs
	for _, f := range fixturesFiles {
		files <- f
	}
	close(files)
	var bytes []byte
	for range fixturesFiles {
		r := <-data
		if len(r) > 0 && len(bytes) > 0 {
			bytes = append(bytes, byte('\n'))
		}
		bytes = append(bytes, r...)
	}
	return bytes
}

func (fl *FileLoader) parse(id int, files <-chan string, data chan<- []byte) {
	for file := range files {
		log.WithFields(log.Fields{
			"worker": id,
			"file":   file,
		}).Debugf("Parsing file : %v", file)
		d, err := readFile(fl.fs, file)
		if err != nil {
			log.WithFields(log.Fields{
				"file":  file,
				"error": err,
			}).Errorf("Could not read file : %v", err)
		}
		data <- d
	}
}

func readFile(fs afero.Fs, file string) ([]byte, error) {

	f, err := afero.ReadFile(fs, file)
	if err != nil {
		return f, err
	}
	return f, nil
}
