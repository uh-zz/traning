package datastructure

import "container/list"

// FYI: https://qiita.com/mattn/items/774280a746c82ee63fc0
type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(v any) {
	s.v.PushBack(v)
}

func (s *Stack) Pop() any {
	e := s.v.Back()
	if e == nil {
		return nil
	}
	return s.v.Remove(e)
}
