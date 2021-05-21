package main

import (
	"log"
	"time"

	"github.com/alecthomas/kong"
	"github.com/guiyomh/charlatan/cmd"
	"github.com/sirupsen/logrus"
)

const (
	Name    = "go-faker-fixtures"
	Author  = "Guillaume Camus"
	Version = "0.1.0"
)

type debugFlag bool

func (d debugFlag) BeforeApply(logger *logrus.Logger) error {
	logger.SetLevel(logrus.DebugLevel)
	return nil
}

var CLI struct {
	Debug debugFlag `help:"Enable debug mode."`

	Host     string `help:"Host Database" default:"127.0.0.1" group:"database"`
	Password string `help:"Database user password" short:"p" required:"" group:"database"`
	User     string `help:"Database username" short:"u" required:"" group:"database"`
	Name     string `help:"Database name" short:"d" required:"" group:"database"`
	Port     int    `help:"Database name" default:"3306" group:"database"`

	Load cmd.LoadCmd `cmd:"" help:"Load fixtures from the path"`

	// Explore cmd.ExploreCmd `cmd:"" help:"Scrap the website"`
	// Replay  cmd.ReplayCmd  `cmd:"" help:"Replay production traffic from datadogs or CSV file"`
}

func main() {
	start := time.Now()
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	ctx := kong.Parse(&CLI, kong.Bind(logger), kong.Name("charlatan"))
	err := ctx.Run()
	ctx.FatalIfErrorf(err)

	elapsed := time.Since(start)
	log.Printf("Excution time : %s", elapsed)
}
