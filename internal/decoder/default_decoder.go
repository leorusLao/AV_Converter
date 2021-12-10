package decoder

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/leorusLao/AV_Converter/common"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

var _ Decoder = (*defaultDecoder)(nil)

type defaultDecoder struct {
	executor  string
	outOption *common.AVOption
}

func NewDefaultDecoder() Decoder {
	return &defaultDecoder{
		executor: common.FFmpeg.String(),
		outOption: &common.AVOption{
			Format: common.FormatWav.String(),
		},
	}
}

func (d *defaultDecoder) Decode(in *common.AVOption) (out *common.AVOption, err error) {
	var args []string

	args = append(args, "-y")
	if in.SampleRate != 0 {
		args = append(args, "-ar", strconv.Itoa(in.SampleRate))
	}
	if in.Channels != 0 {
		args = append(args, "-ac", strconv.Itoa(in.Channels))
	}
	if in.Format != "" {
		args = append(args, "-f", in.Format)
	}
	args = append(args, "-i", in.Path)
	args = append(args, "-f", d.outOption.Format)
	outPath := filepath.Join(os.TempDir(), "conv_"+uuid.New().String()+"."+d.outOption.Format)
	args = append(args, outPath)

	var outBytes bytes.Buffer
	cmd := exec.Command(d.executor, args...)
	cmd.Stdout = &outBytes
	cmd.Stderr = &outBytes
	err = cmd.Run()
	if err != nil {
		return nil, errors.Wrap(errors.Errorf("error: %s, out: %s", err, outBytes.String()), "cmd.Run")
	}

	return &common.AVOption{
		Path: outPath,
	}, nil
}
