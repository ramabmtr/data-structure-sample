package data_structure_sample

import (
	"reflect"
	"testing"
)

func TestQueue_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		createFn func() *Queue
		want     bool
	}{
		{
			name: "queue is empty",
			createFn: func() *Queue {
				return NewQueue(3)
			},
			want: true,
		},
		{
			name: "queue is not empty",
			createFn: func() *Queue {
				a := NewQueue(3)
				_ = a.EnQueue(1)
				return a
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.createFn()
			if got := q.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_IsFull(t *testing.T) {
	tests := []struct {
		name     string
		createFn func() *Queue
		want     bool
	}{
		{
			name: "queue is full",
			createFn: func() *Queue {
				a := NewQueue(3)
				_ = a.EnQueue(1)
				_ = a.EnQueue(2)
				_ = a.EnQueue(3)
				return a
			},
			want: true,
		},
		{
			name: "queue is not full",
			createFn: func() *Queue {
				a := NewQueue(3)
				_ = a.EnQueue(1)
				return a
			},
			want: false,
		},
		{
			name: "queue is empty",
			createFn: func() *Queue {
				return NewQueue(3)
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.createFn()
			if got := q.IsFull(); got != tt.want {
				t.Errorf("IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	tests := []struct {
		name     string
		createFn func() *Queue
		want     any
		wantErr  error
	}{
		{
			name: "empty queue return error",
			createFn: func() *Queue {
				return NewQueue(3)
			},
			want:    nil,
			wantErr: ErrQueueEmpty,
		},
		{
			name: "peek queue with value [1, 2, 3]",
			createFn: func() *Queue {
				a := NewQueue(3)
				_ = a.EnQueue(1)
				_ = a.EnQueue(2)
				_ = a.EnQueue(3)
				return a
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "peek queue with value [1]",
			createFn: func() *Queue {
				a := NewQueue(3)
				_ = a.EnQueue(1)
				return a
			},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.createFn()
			got, err := q.Peek()
			if err != tt.wantErr {
				t.Errorf("Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func TestQueue_DeQueue(t *testing.T) {
	tests := []struct {
		name      string
		createFn  func() *Queue
		want      any
		wantErr   error
		dataAfter []any
	}{
		{
			name: "dequeue empty queue",
			createFn: func() *Queue {
				return NewQueue(3)
			},
			want:      nil,
			wantErr:   ErrQueueEmpty,
			dataAfter: []any{},
		},
		{
			name: "dequeue queue with value [1]",
			createFn: func() *Queue {
				a := NewQueue(3)
				_ = a.EnQueue(1)
				return a
			},
			want:      1,
			wantErr:   nil,
			dataAfter: []any{},
		},
		{
			name: "dequeue queue with value [1, 2]",
			createFn: func() *Queue {
				a := NewQueue(3)
				_ = a.EnQueue(1)
				_ = a.EnQueue(2)
				return a
			},
			want:      1,
			wantErr:   nil,
			dataAfter: []any{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.createFn()
			got, err := q.DeQueue()
			if err != tt.wantErr {
				t.Errorf("DeQueue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeQueue() got = %v, want %v", got, tt.want)
				return
			}
			if !reflect.DeepEqual(q.GetAllData(), tt.dataAfter) {
				t.Errorf("data after DeQueue() = %v, want %v", q.GetAllData(), tt.dataAfter)
				return
			}
		})
	}
}

func TestQueue_EnQueue(t *testing.T) {
	tests := []struct {
		name      string
		createFn  func() *Queue
		data      any
		wantErr   error
		dataAfter []any
	}{
		{
			name: "enqueue to empty queue",
			createFn: func() *Queue {
				return NewQueue(3)
			},
			data:      1,
			wantErr:   nil,
			dataAfter: []any{1},
		},
		{
			name: "enqueue to non empty queue",
			createFn: func() *Queue {
				a := NewQueue(3)
				_ = a.EnQueue(1)
				return a
			},
			data:      2,
			wantErr:   nil,
			dataAfter: []any{1, 2},
		},
		{
			name: "enqueue to full queue",
			createFn: func() *Queue {
				a := NewQueue(1)
				_ = a.EnQueue(1)
				return a
			},
			data:      2,
			wantErr:   ErrQueueFull,
			dataAfter: []any{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.createFn()
			err := q.EnQueue(tt.data)
			if err != tt.wantErr {
				t.Errorf("EnQueue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(q.GetAllData(), tt.dataAfter) {
				t.Errorf("data after DeQueue() = %v, want %v", q.GetAllData(), tt.dataAfter)
				return
			}
		})
	}
}

func TestQueue_Search(t *testing.T) {
	tests := []struct {
		name     string
		createFn func() *Queue
		data     any
		want     int
	}{
		{
			name: "search in empty queue",
			createFn: func() *Queue {
				return NewQueue(3)
			},
			data: 1,
			want: -1,
		},
		{
			name: "search found in queue",
			createFn: func() *Queue {
				a := NewQueue(3)
				_ = a.EnQueue(1)
				_ = a.EnQueue(2)
				_ = a.EnQueue(3)
				return a
			},
			data: 2,
			want: 1,
		},
		{
			name: "search found in queue multiple",
			createFn: func() *Queue {
				a := NewQueue(3)
				_ = a.EnQueue(1)
				_ = a.EnQueue(1)
				_ = a.EnQueue(1)
				return a
			},
			data: 1,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.createFn()
			if got := q.Search(tt.data); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
