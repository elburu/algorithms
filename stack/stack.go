import "fmt"

type Stack struct {
	items []string
}

func (s *Stack) Push(b string) {
	fmt.Println(s.items)
	s.items = append(s.items, b)

}

func (s *Stack) Peek() string {
	stackLen := len(s.items)
	if stackLen > 0 {
		return s.items[stackLen-1]
	}
	return ""
}

func (s *Stack) Pop() string {
	stackLen := len(s.items)
	if stackLen > 0 {
		result := s.items[stackLen-1]
		s.items = s.items[:stackLen-1]
		return result
	}
	return ""
}