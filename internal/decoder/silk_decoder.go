package decoder

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/leorusLao/AV_Converter/common"
	"github.com/pkg/errors"
	"os/exec"
	"path/filepath"
)

var _ Decoder = (*silkDecoder)(nil)

type silkDecoder struct {
	executor  string
	outOption *common.AVOption
}

func NewSilkDecoder() Decoder {
	return &silkDecoder{
		executor: common.SilkDecoder.String(),
		outOption: &common.AVOption{
			Format: common.FormatPcm.String(),
		},
	}
}

func (d *silkDecoder) Decode(in *common.AVOption) (out *common.AVOption, err error) {
	var args []string

	args = append(args, in.Path)
	outPath := filepath.Join("./", "conv_"+uuid.New().String()+"."+d.outOption.Format)
	args = append(args, outPath)
	args = append(args, "-Fs_API", "16000")
	args = append(args, "-quiet")

	var outBytes bytes.Buffer
	cmd := exec.Command(d.executor, args...)
	cmd.Stdout = &outBytes
	cmd.Stderr = &outBytes
	err = cmd.Run()
	if err != nil {
		return nil, errors.Wrap(errors.Errorf("error: %s, out: %s", err, outBytes.String()), "cmd.Run")
	}

	// pcm 需要指定详细参数
	return &common.AVOption{
		Path:       outPath,
		Format:     "s16le",
		SampleRate: 16000,
		Channels:   1,
	}, nil
}
