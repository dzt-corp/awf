package utils

import (
	"encoding/json"
	"math"
	"os/exec"
	"strconv"
)

// GetPeaks executes bbc/audiowaveform and parses the returned JSON data to
// extract the waveform peaks.
// Returns the array of waveform peaks.
func GetPeaks(filepath string, count int, duration int) ([]int, error) {
	pps := int(math.Ceil(float64(count*1e3) / float64(duration)))
	// TODO: Replace command execution with direct interfacing
	cmd := exec.Command(
		"audiowaveform",
		"--input-filename",
		filepath,
		"--pixels-per-second",
		strconv.Itoa(pps),
		"--output-format",
		"json",
	)
	out, err := cmd.Output()
	if err != nil {
		return []int{}, err
	}

	data := struct {
		// the peaks of the waveform; This consists of alternating positive and
		// negative numbers so only the positive values are filtered and normalised.
		Data []int `json:"data"`
	}{}
	err = json.Unmarshal(out, &data)
	if err != nil {
		return []int{}, err
	}

	return data.Data, nil
}
