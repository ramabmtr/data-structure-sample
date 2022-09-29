package data_structure_sample

import (
	"fmt"
	"sync"
)

var (
	ErrQueueEmpty = fmt.Errorf("queue is empty")
	ErrQueueFull  = fmt.Errorf("queue is full of capacity")
)

type Queue struct {
	data []any
	cap  int
	m    sync.Mutex
}

func NewQueue(cap int) *Queue {
	return &Queue{
		data: make([]any, 0),
		cap:  cap,
		m:    sync.Mutex{},
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) IsFull() bool {
	return len(q.data) == q.cap
}

func (q *Queue) EnQueue(data any) error {
	q.m.Lock()
	defer q.m.Unlock()

	if q.IsFull() {
		return ErrQueueFull
	}

	q.data = append(q.data, data)

	return nil
}

func (q *Queue) DeQueue() (any, error) {
	q.m.Lock()
	defer q.m.Unlock()

	if q.IsEmpty() {
		return nil, ErrQueueEmpty
	}

	data := q.data[0]
	q.data = q.data[1:]

	return data, nil
}

func (q *Queue) Peek() (any, error) {
	q.m.Lock()
	defer q.m.Unlock()

	if q.IsEmpty() {
		return nil, ErrQueueEmpty
	}

	return q.data[0], nil
}

func (q *Queue) Search(data any) int {
	for i, datum := range q.data {
		if data == datum {
			return i
		}
	}

	return -1
}

func (q *Queue) GetAllData() []any {
	q.m.Lock()
	defer q.m.Unlock()

	return q.data
}
