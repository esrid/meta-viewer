package logic

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

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
	Stdout   bytes.Buffer
	Stderr   bytes.Buffer
}

type Options func(*ExifTool)

func NewExif(options ...Options) (*ExifTool, error) {
	exiftool := &ExifTool{}

	for _, opt := range options {
		opt(exiftool)
	}
	// argums := strings.Join(exiftool.Options, " ")

	cmd := exec.Command("exiftool", fmt.Sprintf("%s", exiftool.filename))

	cmd.Stdout = &exiftool.Stdout
	cmd.Stderr = &exiftool.Stderr

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return exiftool, nil
}

func Withfilename(filename string) Options {
	return func(et *ExifTool) {
		et.filename = filename
	}
}

func WithOptions(options ...string) Options {
	return func(et *ExifTool) {
		for _, opt := range options {
			et.Options = append(et.Options, opt)
		}
	}
}

func Scanner(Stdout bytes.Buffer) map[string][]string {
	scanOutMap := make(map[string]string)

	scanner := bufio.NewScanner(&Stdout)

	for scanner.Scan() {
		line := scanner.Text()
		astr := strings.Split(line, ":")

		key := strings.TrimSpace(astr[0])
		value := strings.TrimSpace(astr[1])

		if key == exifVersion {
			continue
		}
		scanOutMap[key] = value
	}
	parsedMap := make(map[string][]string)

	for key, value := range scanOutMap {
		if args := containsOptions(value, options); args != "" {
			parsedMap[key] = append(parsedMap[key], value, args)
		} else {
			parsedMap[key] = []string{value}
		}
	}

	return parsedMap
}

func containsOptions(str string, Sslice []string) string {
	var opts []string
	for _, value := range Sslice {
		if exist := strings.Contains(str, value); exist {
			opts = append(opts, value)
		}
	}
	return strings.Join(opts, " ")
}
