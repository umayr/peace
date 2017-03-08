package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/umayr/peace"
)

var (
	pkg     string
	tags    string
	logging bool
)

func init() {
	flag.StringVar(&pkg, "pkg", "", "package name")
	flag.StringVar(&tags, "tags", "", "additional tags")
	flag.BoolVar(&logging, "log", false, "print logs")
}

func main() {
	flag.Parse()

	if pkg == "" {
		log.Fatal("Package name is required")
	}

	if strings.HasPrefix(pkg, "/") {
		log.Fatal("Invalid package name")
	}

	r, err := peace.Do(pkg, tags, logging)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(r)
}
