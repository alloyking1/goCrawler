package main

import (
	"fmt"
	"flag"
	"os"
)

func main () {
	flag.Parse()
	args:= flag.Args()

	if len (args) < 1{
		fmt.Println("Please provide me a page to craw");
		os.Exit(1)
	}

}