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

func READATA(filepath string) (map[string]string, error) {
	cmd := exec.Command("exiftool", fmt.Sprintf("%s", filepath))
	var out, outerr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &outerr

	if err := cmd.Run(); err != nil {
		return map[string]string{}, err
	}

	outputmap := map[string]string{}

	scanner := bufio.NewScanner(&out)

	for scanner.Scan() {
		astr := strings.Split(scanner.Text(), ":")

		if strings.TrimSpace(astr[0]) == exifVersion {
			continue
		}
		key := strings.TrimSpace(astr[0])
		value := strings.TrimSpace(astr[1])

		outputmap[key] = value
	}
	return outputmap, nil
}
