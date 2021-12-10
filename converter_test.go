package main

import (
	"github.com/leorusLao/AV_Converter/common"
	"log"
	"os"
	"testing"
)

var (
	audio  = "./testdata/audio/"
	video  = "./testdata/video/"
	output = "./output/"
)

func init() {
	if err := os.RemoveAll(output); err != nil {
		log.Fatalf("os.RemoveAll error: %s", err)
	}
	if err := os.MkdirAll(output, os.ModePerm); err != nil {
		log.Fatalf("os.RemoveAll error: %s", err)
	}
}

func fileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func TestConvert(t *testing.T) {
	tests := []struct {
		in   string
		outs []string
	}{
		// wav
		{
			in: audio + "16K.wav",
			outs: []string{
				output + "wav.wav",
				output + "wav.mp3",
				output + "wav.flac",
				output + "wav.ogg",
			},
		},
		// aac
		{
			in: audio + "aac.aac",
			outs: []string{
				output + "aac.wav",
				output + "aac.mp3",
				output + "aac.flac",
				output + "aac.ogg",
			},
		},
		// ac3
		{
			in: audio + "ac3.ac3",
			outs: []string{
				output + "ac3.wav",
				output + "ac3.mp3",
				output + "ac3.flac",
				output + "ac3.ogg",
			},
		},
		// adu
		{
			in: audio + "adu.adu",
			outs: []string{
				output + "adu.wav",
				output + "adu.mp3",
				output + "adu.flac",
				output + "adu.ogg",
			},
		},
		// 其它略
	}

	for _, v := range tests {
		inOpt := &common.AVOption{Path: v.in}
		for _, o := range v.outs {
			outOpt := &common.AVOption{Path: o}
			converter := newConverter(inOpt, outOpt)
			if err := converter.Convert(); err != nil {
				t.Fatalf("converter.Convert: %s", err)
			}

			if !fileExist(o) {
				t.Fatalf("convert fail, output not exist!")
			}
		}
	}
}
