package common

type Format string

const (
	FormatWav    Format = "wav"    // audio
	FormatAmr    Format = "amr"    // audio
	FormatAud    Format = "aud"    // audio
	FormatSilk   Format = "silk"   // audio
	FormatSlk    Format = "slk"    // audio
	FormatAc3    Format = "ac3"    // audio
	FormatAiff   Format = "aiff"   // audio
	FormatApe    Format = "ape"    // audio
	FormatMp2    Format = "mp2"    // audio
	FormatMp3    Format = "mp3"    // audio
	FormatWma    Format = "wma"    // audio
	FormatM4a    Format = "m4a"    // audio
	FormatAac    Format = "aac"    // audio
	FormatAdts   Format = "adts"   // audio
	FormatAdu    Format = "adu"    // audio
	FormatHzmv   Format = "hzmv"   // audio
	FormatSdp    Format = "sdp"    // audio
	FormatPcm    Format = "pcm"    // audio
	FormatSpx    Format = "spx"    // audio
	FormatBackup Format = "backup" // audio
	FormatMkv    Format = "mkv"    // video
	FormatMov    Format = "mov"    // video
	FormatMp4    Format = "mp4"    // video
	FormatFlv    Format = "flv"    // video
	FormatM4v    Format = "m4v"    // video
	FormatMpg    Format = "mpg"    // video
	FormatSwf    Format = "swf"    // video
	FormatVob    Format = "vob"    // video
	FormatWmv    Format = "wmv"    // video
	FormatAvi    Format = "avi"    // video
	FormatAsf    Format = "asf"    // video
	Format3gp    Format = "3gp"    // video
)

func (f Format) String() string {
	return string(f)
}