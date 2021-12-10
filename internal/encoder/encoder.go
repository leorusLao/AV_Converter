package encoder

import (
	"github.com/leorusLao/AV_Converter/common"
)

type Encoder interface {
	Encode(*common.AVOption, *common.AVOption) error
}

type NewEncoderFunc func() Encoder

var NewEncoderFuncMap map[common.Format]NewEncoderFunc

func init() {
	NewEncoderFuncMap = make(map[common.Format]NewEncoderFunc)
}

func UseEncoder(format string) Encoder {
	f, ok := NewEncoderFuncMap[common.Format(format)]
	if !ok {
		f = NewDefaultEncoder
	}
	return f()
}
