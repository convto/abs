package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("missing parameter count.\nonly support one argment")
		os.Exit(1)
	}
	path := os.Args[1]
	f, err := os.Open(path)
	f.Close()
	if err != nil {
		fmt.Printf("filepath check error.\n%v\n", err)
		os.Exit(1)
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("filepath convert error.\n%v\n", err)
		os.Exit(1)
	}
	fmt.Println(abs)
}
