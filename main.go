package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/spraylist/generate"
)

const (
	version    = "1.1.1"
	defproduct = "Unknown"
	usage      = `
An utility for weak password list generator.
Usage:
  sprylist -c config.toml
Options:
  -V              show version and exit
  -h              show usage and exit
  -c              config file (default config.toml)
`
)

func main() {
	showVersion := flag.Bool("V", false, "")
	cfgfile := flag.String("c", "config.toml", "Config file")
	flag.Usage = func() {
		fmt.Println(usage)
	}
	flag.Parse()
	if *showVersion {
		fmt.Println("Programm version:", version)
		os.Exit(0)
	}
	dict, err := generate.New(*cfgfile)
	if err != nil {
		log.Fatalln(err)
	}
	dict.MakeDictonary()
	dict.Write()
}
