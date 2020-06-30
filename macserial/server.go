package main

import (
	"errors"
	"os/exec"
	"strings"
)

//SerialInfoItem struct
type SerialInfoItem struct {
	Invalid bool
	Key     string
	Value   interface{}
}

//SerialInfo struct
type SerialInfo struct {
	Country SerialInfoItem
	Year    SerialInfoItem
	Week    SerialInfoItem
	Line    SerialInfoItem
	Model   SerialInfoItem

	Valid string
}

func genSerialInfoItem(item string, out *SerialInfoItem) {
	data := strings.SplitN(item, "-", 1)
	out.Key, out.Value = strings.TrimSpace(data[0]), strings.TrimSpace(data[1])

	if strings.Contains(item, "Unknown") || strings.Contains(item, "-1") {
		out.Invalid = true
	}
}

func parseMacSerialOutput(bdata []byte) (SerialInfo, error) {
	data := string(bdata)
	if strings.Contains(data, "ERROR") {
		return SerialInfo{}, errors.New("Error occured while running macserial")
	}

	winCompat := strings.ReplaceAll(data, "\r\n", "\n")
	lines := strings.Split(winCompat, "\n")

	result := SerialInfo{}
	genSerialInfoItem(lines[0], &result.Country)
	genSerialInfoItem(lines[1], &result.Year)
	genSerialInfoItem(lines[2], &result.Week)
	genSerialInfoItem(lines[3], &result.Line)
	genSerialInfoItem(lines[4], &result.Model)

	return result, nil
}

func main() {
	cmd := exec.Command("/usr/local/bin/macserial", "--info", "SERIAL")
}
