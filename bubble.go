package algos

import (
    "errors"
    "fmt"
)

func BubbleSort(arr []int) ([]int, error) {
    if len(arr) < 1 {
        return arr, errors.New(fmt.Sprintf("'%v' needs at least 1 element", arr))
    }else if len(arr) == 1 {
        return arr, nil
    }
    
    for i:=0; i < len(arr); i++ {
        for j:=0; j < len(arr) - 1 - i; j++ {
            if arr[j] > arr[j + 1]{
                tmp := arr[j]
                arr[j] = arr[j + 1]
                arr[j + 1] = tmp
            }
        }
    }
    return arr, nil
}
