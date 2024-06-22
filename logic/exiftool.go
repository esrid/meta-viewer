package logic

const (
	exifVersion string = "ExifTool Version Number"
)

var options = []string{
	"-a",
	"-b",
	"-c",
	"-charset",
	"-csv",
	"-d",
	"-D",
	"-e",
	"-E",
	"-ex",
	"-ext",
	"-f",
	"-F",
	"-fast",
	"-fast2",
	"-G",
	"-g",
	"-h",
	"-H",
	"-htmlDump",
	"-i",
	"-if",
	"-k",
	"-l",
	"-lang",
	"-L",
	"-listItem",
	"-listx",
	"-m",
	"-n",
	"-p",
	"-P",
	"-password",
	"-progress",
	"-r",
	"-S",
	"-s",
	"-sep",
	"-short",
	"-t",
	"-T",
	"-v",
	"-w",
	"-W",
	"-x",
	"-X",
	"-z",
	"-@",
	"-globalTimeShift",
	"-use",
	"-restore_original",
	"-delete_original",
	"-api",
	"-common_args",
	"-config",
	"-echo",
	"-efile",
	"-execute",
	"-file",
}

type ExifTool struct {
	filename string
	Options  []string
}

type Options func(*ExifTool)

func (et *ExifTool) NewExif(options ...Options) {
}

func Withfilename(filename string) Options {
	return func(et *ExifTool) {
		et.filename = filename
	}
}

func WithOptions(options ...string) Options {
	return func(et *ExifTool) {
		e
	}
}
