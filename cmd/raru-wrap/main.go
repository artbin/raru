package main

import (
	"log"
	"os"

	"github.com/ArtemKulyabin/raru"
)

func main() {
	exer, err := raru.NewExecutor()
	if err != nil {
		log.Printf("raru critical failure: %s", err.Error())
		os.Exit(1)
	}
	if err := exer.Exec(os.Args[0], os.Args[1:]...); err != nil {
		log.Printf("raru critical failure: %s", err.Error())
		os.Exit(1)
	}
}
