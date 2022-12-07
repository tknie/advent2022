package advent

import (
	"errors"
	"sync"
)

// Stack stack creates structure
type Stack struct {
	lock sync.Mutex // you don't have to do this if you don't want thread safety
	s    []interface{}
	Size int
}

// NewStack creates a new stack instance
func NewStack() *Stack {
	return &Stack{sync.Mutex{}, make([]interface{}, 0), 0}
}

// Push push a new element into stack
func (s *Stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
	s.Size++
}

// Low push a new element at last
func (s *Stack) Low(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	xx := make([]interface{}, 0)
	xx = append(xx, v)
	xx = append(xx, s.s...)
	s.s = xx
	s.Size++
}

// Pop pop a new element out of stack. If empty a nil interface is returned. Error is indicating the case
func (s *Stack) Pop() (interface{}, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return nil, errors.New("empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	s.Size--

	return res, nil
}

// Clear Clear the stack
func (s *Stack) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = s.s[:0]
}
