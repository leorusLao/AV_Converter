package main

type Converter struct {
	InputOption  *AVOption
	OutputOption *AVOption
}

type AVOption struct {
	Path       string
	Format     string
	SampleRate int
	Channels   int
	BitDepth   int
}

func newConverter() *Converter {
	return &Converter{}
}
