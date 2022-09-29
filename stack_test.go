package data_structure_sample

import (
	"reflect"
	"testing"
)

func TestStack_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		createFn func() *Stack
		want     bool
	}{
		{
			name: "stack is empty",
			createFn: func() *Stack {
				return NewStack(3)
			},
			want: true,
		},
		{
			name: "stack is not empty",
			createFn: func() *Stack {
				a := NewStack(3)
				_ = a.Push(1)
				return a
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.createFn()
			if got := s.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_IsFull(t *testing.T) {
	tests := []struct {
		name     string
		createFn func() *Stack
		want     bool
	}{
		{
			name: "stack is full",
			createFn: func() *Stack {
				a := NewStack(3)
				_ = a.Push(1)
				_ = a.Push(2)
				_ = a.Push(3)
				return a
			},
			want: true,
		},
		{
			name: "stack is not full",
			createFn: func() *Stack {
				a := NewStack(3)
				_ = a.Push(1)
				return a
			},
			want: false,
		},
		{
			name: "stack is empty",
			createFn: func() *Stack {
				return NewStack(3)
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.createFn()
			if got := s.IsFull(); got != tt.want {
				t.Errorf("IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name     string
		createFn func() *Stack
		want     any
		wantErr  error
	}{
		{
			name: "empty stack return error",
			createFn: func() *Stack {
				return NewStack(3)
			},
			want:    nil,
			wantErr: ErrStackEmpty,
		},
		{
			name: "peek stack with value [1, 2, 3]",
			createFn: func() *Stack {
				a := NewStack(3)
				_ = a.Push(1)
				_ = a.Push(2)
				_ = a.Push(3)
				return a
			},
			want:    3,
			wantErr: nil,
		},
		{
			name: "peek stack with value [1]",
			createFn: func() *Stack {
				a := NewStack(3)
				_ = a.Push(1)
				return a
			},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.createFn()
			got, err := s.Peek()
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

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name      string
		createFn  func() *Stack
		want      any
		wantErr   error
		dataAfter []any
	}{
		{
			name: "pop empty stack",
			createFn: func() *Stack {
				return NewStack(3)
			},
			want:      nil,
			wantErr:   ErrStackEmpty,
			dataAfter: []any{},
		},
		{
			name: "pop stack with value [1]",
			createFn: func() *Stack {
				a := NewStack(3)
				_ = a.Push(1)
				return a
			},
			want:      1,
			wantErr:   nil,
			dataAfter: []any{},
		},
		{
			name: "pop stack with value [1, 2]",
			createFn: func() *Stack {
				a := NewStack(3)
				_ = a.Push(1)
				_ = a.Push(2)
				return a
			},
			want:      2,
			wantErr:   nil,
			dataAfter: []any{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.createFn()
			got, err := s.Pop()
			if err != tt.wantErr {
				t.Errorf("Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
				return
			}
			if !reflect.DeepEqual(s.GetAllData(), tt.dataAfter) {
				t.Errorf("data after DeQueue() = %v, want %v", s.GetAllData(), tt.dataAfter)
				return
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	tests := []struct {
		name      string
		createFn  func() *Stack
		data      any
		wantErr   error
		dataAfter []any
	}{
		{
			name: "push to empty stack",
			createFn: func() *Stack {
				return NewStack(3)
			},
			data:      1,
			wantErr:   nil,
			dataAfter: []any{1},
		},
		{
			name: "push to non empty stack",
			createFn: func() *Stack {
				a := NewStack(3)
				_ = a.Push(1)
				return a
			},
			data:      2,
			wantErr:   nil,
			dataAfter: []any{1, 2},
		},
		{
			name: "push to full stack",
			createFn: func() *Stack {
				a := NewStack(1)
				_ = a.Push(1)
				return a
			},
			data:      2,
			wantErr:   ErrStackFull,
			dataAfter: []any{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.createFn()
			err := s.Push(tt.data)
			if err != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(s.GetAllData(), tt.dataAfter) {
				t.Errorf("data after DeQueue() = %v, want %v", s.GetAllData(), tt.dataAfter)
				return
			}
		})
	}
}

func TestStack_Search(t *testing.T) {
	tests := []struct {
		name     string
		createFn func() *Stack
		data     any
		want     int
	}{
		{
			name: "search in empty stack",
			createFn: func() *Stack {
				return NewStack(3)
			},
			data: 1,
			want: -1,
		},
		{
			name: "search found in stack",
			createFn: func() *Stack {
				a := NewStack(3)
				_ = a.Push(1)
				_ = a.Push(2)
				_ = a.Push(3)
				return a
			},
			data: 2,
			want: 1,
		},
		{
			name: "search found in stack multiple",
			createFn: func() *Stack {
				a := NewStack(3)
				_ = a.Push(1)
				_ = a.Push(1)
				_ = a.Push(1)
				return a
			},
			data: 1,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.createFn()
			if got := s.Search(tt.data); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
