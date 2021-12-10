package decoder

import (
	"github.com/leorusLao/AV_Converter/common"
)

type Decoder interface {
	Decode(*common.AVOption) (*common.AVOption, error)
}

type NewDecoderFunc func() Decoder

var NewDecoderFuncMap map[common.Format]NewDecoderFunc

func init() {
	NewDecoderFuncMap = make(map[common.Format]NewDecoderFunc)
	NewDecoderFuncMap[common.FormatSilk] = NewSilkDecoder
	NewDecoderFuncMap[common.FormatBackup] = NewSilkDecoder
	NewDecoderFuncMap[common.FormatAdu] = NewAduDecoder
	NewDecoderFuncMap[common.FormatHzmv] = NewAduDecoder
	NewDecoderFuncMap[common.FormatSpx] = NewSpxDecoder
}

func UseDecoder(format string) Decoder {
	f, ok := NewDecoderFuncMap[common.Format(format)]
	if !ok {
		f = NewDefaultDecoder
	}
	return f()
}
