package algos

import (
    "testing"
    "reflect"    
)

func TestBubbleSort(t *testing.T){
	arr := []int{9, 3, 7, 4, 69, 420, 42}

	want := []int{3, 4, 7, 9, 42, 69, 420}
	res, err := BubbleSort(arr)
	if !reflect.DeepEqual(want, res)  || err != nil {
		t.Fatal(err)
	}

}
