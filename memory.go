package memory

import (
	"strconv"
	"strings"

	"github.com/ssleert/sfolib"
)

type Ram struct {
	Total     int
	Free      int
	Available int
	SwapTotal int
	SwapFree  int
}

const (
	MemInfo = "/proc/meminfo"
)

// parse value from line in /proc/meminfo
func parseMemValue(v string) int {
	in, _ := strconv.Atoi(strings.Fields(v)[1])
	return in
}

// get Ram struct from /proc/meminfo
func GetRam() (Ram, error) {
	lines, err := sfolib.ReadLines(MemInfo, 16)
	if err != nil {
		return Ram{}, err
	}

	ram := Ram{
		Total:     parseMemValue(lines[0]),
		Free:      parseMemValue(lines[1]),
		Available: parseMemValue(lines[2]),
		SwapTotal: parseMemValue(lines[14]),
		SwapFree:  parseMemValue(lines[15]),
	}

	return ram, nil
}

// get total ram amount
func GetTotalRam() (int, error) {
	line, err := sfolib.ReadFirstLine(MemInfo)
	if err != nil {
		return 0, err
	}

	return parseMemValue(line), nil
}

// get free ram amount
func GetFreeRam() (int, error) {
	line, err := sfolib.ReadLine(MemInfo, 2)
	if err != nil {
		return 0, err
	}

	return parseMemValue(line), nil
}

// get available ram amount
func GetAvalibleRam() (int, error) {
	line, err := sfolib.ReadLine(MemInfo, 3)
	if err != nil {
		return 0, err
	}

	return parseMemValue(line), nil
}

func GetSwapTotal() (int, error) {
	line, err := sfolib.ReadLine(MemInfo, 15)
	if err != nil {
		return 0, err
	}

	return parseMemValue(line), nil
}

func GetSwapFree() (int, error) {
	line, err := sfolib.ReadLine(MemInfo, 16)
	if err != nil {
		return 0, err
	}

	return parseMemValue(line), nil
}
