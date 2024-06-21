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

		outputmap[astr[0]] = astr[1]
	}
	return outputmap, nil
}
