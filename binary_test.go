package algos

import "testing"

func TestBinarySearch(t *testing.T){
	haystack := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	needle := 69
	want := 3
	res, err := BinarySearch(haystack, needle)
	if res != want || err != nil {
		t.Fatal(err)
	}

	needle = 1336
	want = -1
	res, err =BinarySearch(haystack, needle)
	if res != want || err == nil {
		t.Fatal(err)
	}

	needle = 69420
	want = 10
	res, err =BinarySearch(haystack, needle)
	if res != want || err != nil {
		t.Fatal(err)
	}

	needle = 69421
	want = -1
	res, err =BinarySearch(haystack, needle)
	if res != want || err == nil {
		t.Fatal(err)
	}

	needle = 1
	want = 0
	res, err =BinarySearch(haystack, needle)
	if res != want || err != nil {
		t.Fatal(err)
	}

	needle = 0 
	want = -1
	res, err =BinarySearch(haystack, needle)
	if res != want || err == nil {
		t.Fatal(err)
	}

}
