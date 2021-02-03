package adt

//Stack a sample stack
type Stack []interface{}

//Push push an element to the stack
func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

//IsEmpty check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

//Pop pop an element from the stack
func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	index := len(*s) - 1
	item := (*s)[index]
	*s = (*s)[:index]
	return item, true
}
