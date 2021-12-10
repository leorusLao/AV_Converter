package config

import (
	"flag"
	"os"
)

type Config struct {
	Input      string
	Output     string
	Formats    arrayF
	SampleRate arrayAR
	Channels   arrayAC
}

var (
	Conf    *Config
	flagSet *flag.FlagSet
)

func InitConf() (*flag.FlagSet, error) {
	flagSet = flag.NewFlagSet("AV_Converter", flag.ExitOnError)

	flagSet.Bool("version", false, "print version string")

	var formats arrayF
	var sampleRates arrayAR
	var channels arrayAC

	flagSet.String("i", "", "input file")
	flagSet.String("o", "", "output file")
	flagSet.Var(&formats, "f", "format")
	flagSet.Var(&sampleRates, "ar", "sample rate")
	flagSet.Var(&channels, "ac", "channel")

	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		return nil, err
	}

	Conf = &Config{
		Input:      flagSet.Lookup("i").Value.String(),
		Output:     flagSet.Lookup("o").Value.String(),
		Formats:    formats,
		SampleRate: sampleRates,
		Channels:   channels,
	}

	return flagSet, nil
}
