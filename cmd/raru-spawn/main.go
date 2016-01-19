package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/ArtemKulyabin/raru"
)

func usageQuit() {
	fmt.Println("Usage: raru <program> [arguments]")
	fmt.Println("Runs <program> as a random UID and GID (31337-96872).")
	fmt.Println("Recommended: alias raru='env -i PATH=$PATH raru'")
	os.Exit(2)
}

func main() {
	if len(os.Args) == 1 {
		usageQuit()
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := raru.Spawn(cmd); err != nil {
		log.Printf("raru critical failure: %s", err.Error())
		os.Exit(1)
	}
}
