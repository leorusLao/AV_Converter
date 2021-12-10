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
	NewDecoderFuncMap[common.FormatAmr] = NewDefaultDecoder
}
