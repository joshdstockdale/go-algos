package algos

import (
    "errors"
    "fmt"
    "math"
)

func BinarySearch(haystack []int, needle int) (int, error) {
    var lo int = 0
    var hi int = len(haystack)
    for ok := true; ok; ok = lo < hi {
        var mid int = int(math.Floor(float64(lo + (hi - lo)/2)))
        if needle == haystack[mid]{
            return mid, nil
        }else if needle < haystack[mid] {
            hi = mid
        } else {
            lo = mid + 1
        }
    }

    return -1, errors.New(fmt.Sprintf("'%v' not found in '%v'", needle, haystack))
}
