package main

import "github.com/leorusLao/AV_Converter/common"

type Converter struct {
	InputOption  *common.AVOption
	OutputOption *common.AVOption
}

func newConverter() *Converter {
	return &Converter{}
}
