package utils

type Stack struct {
	items []string
}

func (s *Stack) Push(c string) {
	s.items = append(s.items, c)
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return 
	}
	s.items = s.items[:len(s.items) - 1]
}

func (s *Stack) IsEmpty() (bool) {
	return len(s.items) == 0
}
