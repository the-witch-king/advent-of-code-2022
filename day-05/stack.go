package dayfive

type stack struct {
	crates []rune
}

// Multi item operations
func (s *stack) popMany(amount int) []rune {
	if len(s.crates) <= amount {
		r := s.crates
		s.crates = []rune{}
		return r
	}

	r := s.crates[len(s.crates) - amount:]
	s.crates = s.crates[:len(s.crates) - amount]
	return r
}
func (s *stack) pushMany(c []rune) {
	s.crates = append(s.crates, c...)
}

// Single item operations
func (s *stack) unshift(c rune) {
	s.crates = append([]rune{c}, s.crates...)
}
func (s *stack) push(c rune) {
	s.crates = append(s.crates, c)
}
func (s *stack) pop() rune {

	var empty rune
	if len(s.crates) == 0 {
		return empty
	}

	if len(s.crates) == 1 {
		c := s.crates[0]
		s.crates = []rune{}

		return c
	}

	c := s.crates[len(s.crates)-1]
	s.crates = s.crates[:len(s.crates)-1]

	return c
}

// Printing & Visibility
func (s stack) peek() rune {
	if len(s.crates) == 0 {
		return ' '
	}
	return s.crates[len(s.crates)-1]
}
func (s stack) print() string {
	r := ""
	for _, c := range s.crates {
		r += string(c)
	}
	return r
}
