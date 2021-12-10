package common

type Tool string

const (
	FFmpeg      Tool = "./tools/ffmpeg"
	SilkDecoder Tool = "./tools/silk_v3_decoder"
)

func (t Tool) String() string {
	return string(t)
}
