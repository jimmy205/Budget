package main

import (
	"log"
	"time"
)

func main() {
	s, _ := time.Parse("2006-01-02", "2021-01-30")
	// e, _ := time.Parse("2006-01-02", "2021-02-01")

	log.Println(s.Day())

	// log.Println(e.Sub(s).Hours() / 24)
}
