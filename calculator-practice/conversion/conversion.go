package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))

	for stringIndex, str := range strings {
		floatPrice, err := strconv.ParseFloat(str, 64)

		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}

		floats[stringIndex] = floatPrice
	}

	return floats, nil
}
