package utils

import "math"

// filterPositive filters out the negative values from the given array.
func filterPositive(peaks []int) []int {
	posPeaks := make([]int, len(peaks)/2)
	for idx, peak := range peaks {
		if idx%2 == 1 { // odd peaks are positive
			posPeaks[idx/2] = peak
		}
	}
	return posPeaks
}

// getMax finds the largest number in the array.
func getMax(peaks []int) int {
	max := peaks[0]
	for _, peak := range peaks {
		if peak > max {
			max = peak
		}
	}
	return max
}

// roundOff rounds the given float to a fixed number of decimal points.
func roundOff(number float64) float64 {
	ROUNDING_F := 1e5
	number = float64(int(math.Round(number*ROUNDING_F))) / ROUNDING_F
	return number
}

// scaleDown divides all elements in the array by the largest converting them
// into an array of values in the rage [0, 1].
func scaleDown(peaks []int) []float64 {
	scaledPeaks := make([]float64, len(peaks))
	max := getMax(peaks)
	for idx, peak := range peaks {
		scaledPeak := float64(peak) / float64(max)
		scaledPeak = roundOff(scaledPeak)
		scaledPeaks[idx] = scaledPeak
	}
	return scaledPeaks
}

// Normalise takes the bbc/audiowaveform peaks and normalises them into an array
// of fractions in the range [0, 1].
// Returns the normalised form of the given array.
func Normalise(peaks []int) []float64 {
	filteredPeaks := filterPositive(peaks)
	scaledPeaks := scaleDown(filteredPeaks)
	return scaledPeaks
}
