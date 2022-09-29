package data_structure_sample

import (
	"fmt"
	"sync"
)

var (
	ErrStackEmpty = fmt.Errorf("stack is empty")
	ErrStackFull  = fmt.Errorf("stack is full of capacity")
)

type Stack struct {
	data []any
	cap  int
	m    sync.Mutex
}

func NewStack(cap int) *Stack {
	return &Stack{
		data: make([]any, 0),
		cap:  cap,
		m:    sync.Mutex{},
	}
}

func (s *Stack) topIndex() int {
	return len(s.data) - 1
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) IsFull() bool {
	return len(s.data) == s.cap
}

func (s *Stack) Push(data any) error {
	s.m.Lock()
	defer s.m.Unlock()

	if s.IsFull() {
		return ErrStackFull
	}

	s.data = append(s.data, data)

	return nil
}

func (s *Stack) Pop() (any, error) {
	s.m.Lock()
	defer s.m.Unlock()

	if s.IsEmpty() {
		return nil, ErrStackEmpty
	}

	data := s.data[s.topIndex()]
	s.data = s.data[:s.topIndex()]

	return data, nil
}

func (s *Stack) Peek() (any, error) {
	s.m.Lock()
	defer s.m.Unlock()

	if s.IsEmpty() {
		return nil, ErrStackEmpty
	}

	return s.data[s.topIndex()], nil
}

func (s *Stack) Search(data any) int {
	for i, datum := range s.data {
		if data == datum {
			return i
		}
	}

	return -1
}

func (s *Stack) GetAllData() []any {
	s.m.Lock()
	defer s.m.Unlock()

	return s.data
}
