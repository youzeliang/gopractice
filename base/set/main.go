package set

import "sync"

type inter interface {
}

type Set struct {
	m map[inter]bool
	sync.RWMutex
}

func New() *Set {
	return &Set{
		m: make(map[inter]bool),
	}
}

func (s *Set) Add(i inter) {
	s.Lock()
	defer s.Unlock()
	s.m[i] = true
}
