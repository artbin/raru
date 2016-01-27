package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ArtemKulyabin/yax/osx"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	path, err := osx.Executable()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("executable: %s, dir: %s, uid: %d, gid: %d\n", path, wd, os.Getuid(), os.Getgid())
}
