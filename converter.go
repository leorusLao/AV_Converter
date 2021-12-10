package main

import (
	"github.com/leorusLao/AV_Converter/common"
	"github.com/leorusLao/AV_Converter/internal/decoder"
	"github.com/leorusLao/AV_Converter/internal/encoder"
	"github.com/pkg/errors"
	"os"
)

type Converter struct {
	InputOption  *common.AVOption
	OutputOption *common.AVOption
	decoder      decoder.Decoder
	encoder      encoder.Encoder
}

func newConverter(in *common.AVOption, out *common.AVOption) *Converter {
	return &Converter{
		InputOption:  in,
		OutputOption: out,
		decoder:      decoder.UseDecoder(in.GetFormat()),
		encoder:      encoder.UseEncoder(out.GetFormat()),
	}
}

func (c *Converter) Convert() error {
	middle, err := c.decoder.Decode(c.InputOption)
	defer func() {
		if middle != nil {
			_ = os.Remove(middle.Path)
		}
	}()

	if err != nil {
		return errors.Wrap(err, "c.decoder.Decode")
	}

	if err := c.encoder.Encode(middle, c.OutputOption); err != nil {
		return errors.Wrap(err, "c.encoder.Encode")
	}

	return nil
}
