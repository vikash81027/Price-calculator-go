package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) (*[]float64, error) {
	floats := make([]float64, len(strings))

	for strIdx, string := range strings {
		floatPrice, error := strconv.ParseFloat(string, 64)

		if error != nil {
			return nil, errors.New("Failed to convert string to float!")
		}

		floats[strIdx] = floatPrice
	}

	return &floats, nil
}