package utils

import "fmt"

// FormatNumberWithSuffix formats large numbers by adding appropriate magnitude.
// Suffixes are (billion, million, thousand) for better readability
func FormatNumberWithSuffix(value int) string {
	var out string

	switch {
	// value >= 1 billion
	case value >= 1_000_000_000:
		out = fmt.Sprintf("%.2f billion", float64(value)/1_000_000_000)
	// value >= 1 million
	case value >= 1_000_000:
		out = fmt.Sprintf("%.2f million", float64(value)/1_000_000)
	// value >= 1 thousand
	case value >= 1_000:
		out = fmt.Sprintf("%.2f thousand", float64(value)/1_000)
	// for smaller values just return the number
	default:
		out = fmt.Sprintf("%d", value)
	}

	return out
}
