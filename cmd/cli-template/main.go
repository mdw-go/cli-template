package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var Version = "dev"

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	flags := flag.NewFlagSet(fmt.Sprintf("%s @ %s", filepath.Base(os.Args[0]), Version), flag.ExitOnError)
	_ = flags.Parse(os.Args[1:])

	fmt.Println(flags.Name())
}
