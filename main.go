package main

import (
	"log"
	"time"
)

const (
	Name    = "go-faker-fixtures"
	Author  = "Guillaume Camus"
	Version = "0.1.0"
)

func main() {
	start := time.Now()
	// cmd.Execute()
	elapsed := time.Since(start)
	log.Printf("Excution time : %s", elapsed)
}
