package main

import (
	"log"
	"time"
)

func main() {

	s, _ := time.Parse("2006-01-02", "2021-01-01")
	e, _ := time.Parse("2006-01-02", "2020-01-01")

	log.Println(e.Before(s))
}
