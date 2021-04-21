package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mcoops/jargo"
)

func main() {
	filename := flag.String("jar", "", "The Jar file to scan")

	flag.Parse()

	if _, err := os.Stat(*filename); os.IsNotExist(err) {
		log.Fatalf("Valid file must be specified with -jar, not found: %s", *filename)
	}

	jar, err := jargo.GetJarInfo(*filename)

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range jar.Files {
		fmt.Println(f)
	}

}
