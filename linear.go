package algos

import (
	"errors"
	"fmt"
)

func LinearSearch(haystack []int, needle int) (int, error) {
	for i, h := range haystack {
		if h == needle {
			return i, nil
		}
	}
	return -1, errors.New(fmt.Sprintf("'%v' not found in '%v'", needle, haystack))
}
