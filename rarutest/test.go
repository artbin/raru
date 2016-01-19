package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("uid: %d, gid: %d\n", os.Getuid(), os.Getgid())
}
