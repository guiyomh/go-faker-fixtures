package cmd

import (
	"github.com/alecthomas/repr"
	"github.com/guiyomh/charlatan/pkg/fixture"
	"github.com/sirupsen/logrus"
)

type LoadCmd struct {
	Path string `arg:"" name:"path" help:"Paths of fixtures." type:"path"`
}

func (l *LoadCmd) Run(logger *logrus.Logger) error {
	ld := fixture.NewLoader(logger)
	data, err := ld.Load(l.Path)
	if err != nil {
		return err
	}
	rs := fixture.GenerateRecords(data)
	repr.Print(rs)
	return nil
}
