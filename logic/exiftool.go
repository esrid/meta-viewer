package logic

import (
	"bufio"
	"bytes"
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
	options []string
}

func (e *ExifTool) New(filename string) (bytes.Buffer, bytes.Buffer, error) {
	cmd := exec.Command("exiftool", filename)
	var out, outErr bytes.Buffer
	cmd.Stderr = &outErr
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return bytes.Buffer{}, bytes.Buffer{}, err
	}
	return out, out, nil
}

func (e *ExifTool) scanner(out bytes.Buffer) map[string]string {
	outputMap := make(map[string]string)
	scanner := bufio.NewScanner(&out)
	if scanner.Scan() {
		line := scanner.Text()
		astr := strings.Split(line, ":")

		key := astr[0]
		value := astr[1]

		outputMap[key] = value
	}
	return outputMap
}

func containsoptions(str string, options []string) string {
	for _, option := range options {
		if strings.Contains(str, option) {
			return option
		}
	}
	return ""
}

func (e *ExifTool) parser(m map[string]string, options []string) map[string][]string {
	outputMap := make(map[string][]string)
	for k, v := range m {
		if w := containsoptions(v, options); w != "" {
			outputMap[k] = append(outputMap[k], v, w)
		}
	}
	return outputMap
}
