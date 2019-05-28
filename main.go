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

package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/guiyomh/charlatan/contract"
	"github.com/guiyomh/charlatan/loader"
	"github.com/guiyomh/charlatan/normalizer"
	"github.com/guiyomh/charlatan/parser"
	"github.com/sarulabs/di"

	"github.com/spf13/afero"

	log "github.com/sirupsen/logrus"
)

// App is a dependency container that contain service
var App di.Container

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		DisableColors: false,
	})
	var err error
	App, err = createApp()
	if err != nil {
		log.Panicf("Could not initialize the app: %v", err)
		os.Exit(1)
	}
}

func handleDicCast(service string) {
	log.Errorf("Could not retrieve the service '%v'", service)
}

func main() {
	defer func() {
		App.Delete()
	}()
	dir := "./data"
	loc, ok := App.Get("app.locator").(loader.Locatorer)
	if !ok {
		handleDicCast("app.locator")
		os.Exit(1)
	}
	files, err := loc.LocateFiles(dir)
	if err != nil {
		log.Errorf("Error : Could not find fixtures files in %v\n", dir)
		os.Exit(1)
	}
	ld, ok := App.Get("app.fileloader").(loader.Loader)
	if !ok {
		handleDicCast("app.fileloader")
		os.Exit(1)
	}
	data := ld.Load(files)
	p, ok := App.Get("app.parser").(parser.Parser)
	if !ok {
		handleDicCast("app.parser")
		os.Exit(1)
	}
	d, err := p.Parse(data)
	if err != nil {
		log.Errorf("Parsing fixture : %v", err)
		os.Exit(1)
	}

	denorm, ok := App.Get("app.normalizer.registry").(contract.Register)
	if !ok {
		handleDicCast("app.normalizer.registry")
		os.Exit(1)
	}

	fmt.Printf("struct Data : %v\n", d)
	spew.Dump(d)

	var bag contract.Bager
	for table, data := range *d {
		tableBag, err := denorm.Denormalize(table, data)
		if err != nil {
			log.WithField("table", table).Errorf("Could not denormalize the table '%s'", table)
		}
		if bag != nil {
			bag, err = bag.MergeWith(tableBag)
			if err != nil {
				log.WithField("table", table).Errorf("Could not merge fixture bag of the table '%s'", table)
			}
		} else {
			bag = tableBag
		}
	}
	fmt.Println("------------------")
	fmt.Println(bag)
}

var services []di.Def = []di.Def{
	{
		Name:  "app.fs",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return afero.NewOsFs(), nil
		},
	},
}

func createApp() (di.Container, error) {

	builder, err := di.NewBuilder(di.App)
	if err != nil {
		log.Errorf("Could not instantiate the DIC: %v", err)
		return nil, err
	}
	err = builder.Add(services...)
	if err != nil {
		log.Errorf("Could not build DIC services: %v", err)
		return nil, err
	}
	err = builder.Add(loader.Services...)
	if err != nil {
		log.Errorf("Could not build DIC loader services: %v", err)
		return nil, err
	}
	err = builder.Add(parser.Services...)
	if err != nil {
		log.Errorf("Could not build DIC parser services: %v", err)
		return nil, err
	}
	err = builder.Add(normalizer.Services...)
	if err != nil {
		log.Errorf("Could not build DIC parser services: %v", err)
		return nil, err
	}
	return builder.Build(), nil
}
