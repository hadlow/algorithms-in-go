package main

type Queue struct {
	Items []int
}

func NewQueue() *Queue {
	queue := make([]int, 0)

	return &Queue {
		Items: queue,
	}
}

func (q *Queue) Enqueue(item int) {
	q.Items = append(q.Items, item)
}

func (q *Queue) Dequeue() {
	if !q.IsEmpty() {
		q.Items = q.Items[1:]
	}
}

func (q *Queue) Peek() int {
	if q.IsEmpty() {
		return 0
	}

	return q.Items[0]
}

func (q *Queue) IsEmpty() bool {
	if len(q.Items) == 0 {
		return true
	}

	return false
}
