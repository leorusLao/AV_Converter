package common

import "path/filepath"

type AVOption struct {
	Path       string
	Format     string
	SampleRate int
	Channels   int
}

func (avo *AVOption) GetFormat() string {
	if avo.Format != "" {
		return avo.Format
	}
	ext := filepath.Ext(avo.Path)
	if ext == "" {
		return ""
	}
	return ext[1:]
}
