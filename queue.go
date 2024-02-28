package algos

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	length int
	head   *Node
	tail   *Node
}

func NewQueue() *Queue {
	return &Queue{length: 0}
}

func (q *Queue) Enqueue(val int) error {
	node := &Node{value: val}
	if q.length == 0 {
		q.tail = node
		q.head = node
	}

	q.length++
	q.tail.next = node
	q.tail = node
	return nil
}

func (q *Queue) Deque() (int, error) {
	if q.head == nil {
		return -1, nil
	}
	q.length--
	headVal := q.head.value
	q.head = q.head.next

	return headVal, nil
}

func (q *Queue) Peek() (int, error) {
	if q.head == nil {
		return -1, nil
	}
	return q.head.value, nil
}
