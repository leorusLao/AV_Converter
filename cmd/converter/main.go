package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	Version         = ""
	BuildTime       = ""
)

func main() {
	flagSet, err := initConf()
	if err != nil {
		log.Fatalf("config.InitConf err: %s", err)
	}

	if flagSet.Lookup("version").Value.(flag.Getter).Get().(bool) {
		printVersion()
		os.Exit(0)
	}

	fmt.Println(Conf)

	fmt.Println(Conf.Formats.Get(0))
	fmt.Println(Conf.Formats.Get(1))
}

func printVersion() {
	fmt.Printf("\nAV_Convert: %s\n", Version)
	fmt.Printf("\nVersion: %s\n", Version)
	fmt.Printf("\nBuildTime: %s\n\n", BuildTime)
}