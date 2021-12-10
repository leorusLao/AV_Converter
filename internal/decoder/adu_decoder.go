package decoder

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/leorusLao/AV_Converter/common"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"path/filepath"
)

var _ Decoder = (*aduDecoder)(nil)

type aduDecoder struct {
	executor  string
	outOption *common.AVOption
}

func NewAduDecoder() Decoder {
	return &aduDecoder{
		executor: common.AduConvert.String(),
		outOption: &common.AVOption{
			Format: common.FormatWav.String(),
		},
	}
}

func (d *aduDecoder) Decode(in *common.AVOption) (out *common.AVOption, err error) {
	var args []string

	args = append(args, in.Path)
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
		Path:   outPath,
	}, nil
}
