package common

type Tool string

const (
	FFmpeg      Tool = "./tools/ffmpeg"
	SilkDecoder Tool = "./tools/silk_v3_decoder"
	AduConvert  Tool = "./tools/adu_convert"
	SpxDecoder  Tool = "./tools/speexdec"
)

func (t Tool) String() string {
	return string(t)
}
