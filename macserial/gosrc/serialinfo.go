package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

//SerialInfoItem struct
type SerialInfoItem struct {
	Invalid bool        `json:"Invalid"`
	Value   interface{} `json:"Value"`
}

//SerialInfo struct
type SerialInfo struct {
	Country SerialInfoItem `json:"Country"`
	Year    SerialInfoItem `json:"Year"`
	Week    SerialInfoItem `json:"Week"`
	Line    SerialInfoItem `json:"Line"`
	Model   SerialInfoItem `json:"Model"`

	Valid string `json:"Valid"`
}

func genSerialInfoItem(item string, out *SerialInfoItem) {
	data := strings.SplitN(strings.TrimSpace(item), "-", 2)
	out.Value = strings.TrimSpace(data[1])

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

//GetSerialInfo func
func GetSerialInfo(serial string) (SerialInfo, error) {
	cmd := exec.Command("/usr/bin/macserial", "--info", serial)
	data, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error while running macserial: %s\n", err.Error())
		return SerialInfo{}, err
	}

	return parseMacSerialOutput(data)
}