package main

import (
	"flag"
)

type value struct {
	io uint8 // 0: input, 1: output
	v  string
}

// 格式
type arrayF []value

func (i *arrayF) String() string {
	return "formats"
}

func (i *arrayF) Set(v string) error {
	var io uint8
	flagSet.Visit(func(flag *flag.Flag) {
		if flag.Name == "i" {
			io = 1
		}
	})

	*i = append(*i, value{io, v})
	return nil
}

func (i *arrayF) Get(io uint8) string {
	for k := len(*i)-1; k >= 0; k-- {
		if (*i)[k].io == io {
			return (*i)[k].v
		}
	}
	return ""
}

// 采样率
type arrayAR []value

func (i *arrayAR) String() string {
	return "sample rate"
}

func (i *arrayAR) Set(v string) error {
	var io uint8
	flagSet.Visit(func(flag *flag.Flag) {
		if flag.Name == "i" {
			io = 1
		}
	})

	*i = append(*i, value{io, v})
	return nil
}

func (i *arrayAR) Get(io uint8) string {
	for k := len(*i)-1; k >= 0; k-- {
		if (*i)[k].io == io {
			return (*i)[k].v
		}
	}
	return ""
}

// 声道
type arrayAC []value

func (i *arrayAC) String() string {
	return "channels"
}

func (i *arrayAC) Set(v string) error {
	var io uint8
	flagSet.Visit(func(flag *flag.Flag) {
		if flag.Name == "i" {
			io = 1
		}
	})

	*i = append(*i, value{io, v})
	return nil
}

func (i *arrayAC) Get(io uint8) string {
	for k := len(*i)-1; k >= 0; k-- {
		if (*i)[k].io == io {
			return (*i)[k].v
		}
	}
	return ""
}