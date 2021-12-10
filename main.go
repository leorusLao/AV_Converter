package main

import (
	"flag"
	"fmt"
	"github.com/leorusLao/AV_Converter/config"
	"log"
	"os"
)

var (
	Version         = ""
	BuildTime       = ""
)

func main() {
	flagSet, err := config.InitConf()
	if err != nil {
		log.Fatalf("config.InitConf err: %s", err)
	}

	if flagSet.Lookup("version").Value.(flag.Getter).Get().(bool) {
		printVersion()
		os.Exit(0)
	}

	fmt.Println(config.Conf)

	fmt.Println(config.Conf.Formats.Get(0))
	fmt.Println(config.Conf.Formats.Get(1))
}

func printVersion() {
	fmt.Printf("\nAV_Convert: %s\n", Version)
	fmt.Printf("\nVersion: %s\n", Version)
	fmt.Printf("\nBuildTime: %s\n\n", BuildTime)
}