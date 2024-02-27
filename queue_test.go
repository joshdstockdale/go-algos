package algos

import (
	"testing"
)

func TestQueue(t *testing.T) {

	q := NewQueue()

	err := q.Enqueue(5)
	err = q.Enqueue(7)
	err = q.Enqueue(9)

	want := 5
	d, err := q.Deque()
	if d != want || err != nil {
		t.Fatalf("Deque not correct. Wanted %d, got %d, err: %s", want, d, err)
	}
	want = 2
	if q.length != want {
		t.Fatalf("Length not correct. Wanted %d, got %d", want, q.length)
	}

	err = q.Enqueue(11)
	if err != nil {
		t.Fatal("Enqueue error", err)
	}

	want = 7
	d, err = q.Deque()
	if d != want || err != nil {
		t.Fatalf("Deque not correct. Wanted %d, got %d, err: %s", want, d, err)
	}

	want = 9
	d, err = q.Deque()
	if d != want || err != nil {
		t.Fatalf("Deque not correct. Wanted %d, got %d, err: %s", want, d, err)
	}

	want = 11
	p, err := q.Peek()
	if p != want || err != nil {
		t.Fatalf("Peek not correct. Wanted %d, got %d, err: %s", want, p, err)
	}

	want = 11
	d, err = q.Deque()
	if d != want || err != nil {
		t.Fatalf("Deque not correct. Wanted %d, got %d, err: %s", want, d, err)
	}

	want = -1
	d, err = q.Deque()
	if d != want || err != nil {
		t.Fatalf("Deque not correct. Wanted %d, got %d, err: %s", want, d, err)
	}

	want = 0
	if q.length != want {
		t.Fatalf("Length not correct. Wanted %d, got %d", want, q.length)
	}

	err = q.Enqueue(21)
	if err != nil {
		t.Fatal("Enqueue error", err)
	}
	want = 1
	if q.length != want {
		t.Fatalf("Length not correct. Wanted %d, got %d", want, q.length)
	}
	want = 21
	p, err = q.Peek()
	if p != want || err != nil {
		t.Fatalf("PEEK: %+v", q)
		// t.Fatalf("Peek not correct. Wanted %d, got %d, err: %s", want, p, err)
	}
}
