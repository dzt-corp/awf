package models

// Request stores the information contained in the request received by the API.
type Request struct {
	Identifier string `json:"identifier"`
	Url        string `json:"url"`

	// the span of the track in milliseconds
	Duration int `json:"duration"`

	// the set of number of peaks to generate; By default the output will only
	// contain one set of ~1k peaks. The minimum peak count is ~1 peak/s.
	Counts []int `json:"counts"`
}

// PeakSet stores one set of waveform points for the audio track. The higher the
// length, the more fidelity of the peak set.
type PeakSet struct {
	Length int       `json:"length"`
	Peaks  []float64 `json:"peaks"`
}
