package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64

	for _, stringValue := range strings {
		floatValue, err := strconv.ParseFloat(stringValue, 64)

		if err != nil {
			return nil, errors.New("Converting price to float failed")

		}

		floats = append(floats, floatValue)
	}

	return floats, nil
}
