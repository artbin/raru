package main

import (
	"log"
	"os"

	"github.com/ArtemKulyabin/raru"
)

func main() {
	if err := raru.Exec(os.Args[0], os.Args[1:]...); err != nil {
		log.Printf("raru critical failure: %s", err.Error())
		os.Exit(1)
	}
}
