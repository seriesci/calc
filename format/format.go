package format

import (
	"strconv"
)

// Print returns value with given number of decimal places.
func Print(value int, decimals int) string {
	return strconv.FormatFloat(float64(value), 'f', decimals, 64)
}
