package encoder

import (
	"bytes"
	"github.com/leorusLao/AV_Converter/common"
	"github.com/pkg/errors"
	"os/exec"
	"strconv"
)

var _ Encoder = (*defaultEncoder)(nil)

type defaultEncoder struct {
	executor string
}

func NewDefaultEncoder() Encoder {
	return &defaultEncoder{
		executor: common.FFmpeg.String(),
	}
}

func (d *defaultEncoder) Encode(in *common.AVOption, out *common.AVOption) (err error) {
	var args []string

	args = append(args, "-y")

	// input
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

	// output
	if out.SampleRate != 0 {
		args = append(args, "-ar", strconv.Itoa(out.SampleRate))
	}
	if out.Channels != 0 {
		args = append(args, "-ac", strconv.Itoa(out.Channels))
	}
	if out.Format != "" {
		args = append(args, "-f", out.Format)
	}
	args = append(args, out.Path)

	var outBytes bytes.Buffer
	cmd := exec.Command(d.executor, args...)
	cmd.Stdout = &outBytes
	cmd.Stderr = &outBytes
	err = cmd.Run()
	if err != nil {
		return errors.Wrap(errors.Errorf("error: %s, out: %s", err, outBytes.String()), "cmd.Run")
	}

	return nil
}
