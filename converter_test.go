package main

import (
	"github.com/leorusLao/AV_Converter/common"
	"log"
	"os"
	"path/filepath"
	"testing"
)

const (
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

type testData struct {
	in   string
	outs []string
}

func getData(in string) testData {
	format := filepath.Ext(in)[1:]
	return testData{
		in: in,
		outs: []string{
			output + format + ".wav",
			output + format + ".mp3",
			output + format + ".flac",
			output + format + ".ogg",
		}}
}

func TestConvert(t *testing.T) {
	tests := []testData{
		// audio
		getData(audio + "16K.wav"),
		getData(audio + "aac.aac"),
		getData(audio + "ac3.ac3"),
		getData(audio + "adu.adu"),
		getData(audio + "aiff.aiff"),
		getData(audio + "amr.amr"),
		getData(audio + "ape.ape"),
		getData(audio + "aud.aud"),
		getData(audio + "backup.backup"),
		getData(audio + "flac.flac"),
		getData(audio + "hzmv.hzmv"),
		getData(audio + "m4a.m4a"),
		getData(audio + "mp2.mp2"),
		getData(audio + "mp3.mp3"),
		getData(audio + "pcm.pcm"),
		getData(audio + "silk.silk"),
		getData(audio + "slk.slk"),
		getData(audio + "spx.spx"),
		getData(audio + "vox.vox"),
		getData(audio + "wavf.wavf"),
		getData(audio + "wma.wma"),
		// video
		getData(video + "3gp.3gp"),
		getData(video + "asf.asf"),
		getData(video + "avi.avi"),
		getData(video + "flv.flv"),
		getData(video + "m4v.m4v"),
		getData(video + "mkv.mkv"),
		getData(video + "mov.mov"),
		getData(video + "mp4.mp4"),
		getData(video + "mpg.mpg"),
		getData(video + "swf.swf"),
		getData(video + "vob.vob"),
		getData(video + "wmv.wmv"),
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
