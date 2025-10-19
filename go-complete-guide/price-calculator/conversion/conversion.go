package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	var floatValues []float64

	for _, stringValue := range strings {
		floatValue, err := strconv.ParseFloat(stringValue, 64)

		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}
		floatValues = append(floatValues, floatValue)
	}

	return floatValues, nil
}
