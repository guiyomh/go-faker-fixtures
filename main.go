package main

import (
	"log"
	"time"

	"github.com/guiyomh/charlatan/cmd"
)

const (
	Name    = "Charlatan"
	Author  = "Guillaume Camus"
	Version = "0.0.1"
)

func main() {
	start := time.Now()
	cmd.Execute()
	elapsed := time.Since(start)
	log.Printf("Excution time : %s", elapsed)
}
