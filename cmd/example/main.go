package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var Version = "dev"

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	flags := flag.NewFlagSet(fmt.Sprintf("%s @ %s", os.Args[0], Version), flag.ExitOnError)
	_ = flags.Parse(os.Args[1:])
}
