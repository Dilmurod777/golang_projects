package main

type IntStack struct {
	values []int
}

func (s *IntStack) Push(i int) {
	s.values = append(s.values, i)
}

func (s *IntStack) Pop() {
	s.values = s.values[:len(s.values)-1]
}

func main() {
	s := IntStack{}

	s.Push(5)
	s.Push(6)
	s.Push(7)
	s.Push(8)
	s.Push(9)

	s.Pop()
	s.Pop()
	s.Pop()
}
