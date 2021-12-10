package common

type Tool string

const (
	FFmpeg Tool = "./tools/ffmpeg"
)

func (t Tool) String() string {
	return string(t)
}
