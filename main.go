package main

import (
	"flag"
	"fmt"
	"github.com/leorusLao/AV_Converter/common"
	"github.com/leorusLao/AV_Converter/config"
	"log"
	"os"
	"strconv"
)

var (
	Version   = ""
	BuildTime = ""
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

	converter := newConverter(parseOption())
	if err := converter.Convert(); err != nil {
		log.Fatalf("convert fail: %s", err)
	}
}

func printVersion() {
	fmt.Printf("\nAV_Convert: %s\n", Version)
	fmt.Printf("\nVersion: %s\n", Version)
	fmt.Printf("\nBuildTime: %s\n\n", BuildTime)
}

func parseOption() (*common.AVOption, *common.AVOption) {
	conf := config.Conf

	if conf.Input == "" {
		log.Fatalf("-i is require parameter.")
	}

	if conf.Output == "" {
		log.Fatalf("-o is require parameter.")
	}

	sampleRate, err := strconv.Atoi(conf.SampleRate.Get(0))
	if err != nil {
		log.Fatalf("input sample rate illegal")
	}
	channel, err := strconv.Atoi(conf.Channels.Get(0))
	if err != nil {
		log.Fatalf("input channel illegal")
	}
	inOpt := &common.AVOption{
		Path:       conf.Input,
		Format:     conf.Formats.Get(0),
		SampleRate: sampleRate,
		Channels:   channel,
	}

	sampleRate, err = strconv.Atoi(conf.SampleRate.Get(1))
	if err != nil {
		log.Fatalf("input sample rate illegal")
	}
	channel, err = strconv.Atoi(conf.Channels.Get(1))
	if err != nil {
		log.Fatalf("input channel illegal")
	}
	outOpt := &common.AVOption{
		Path:       conf.Input,
		Format:     conf.Formats.Get(0),
		SampleRate: sampleRate,
		Channels:   channel,
	}

	return inOpt, outOpt
}
