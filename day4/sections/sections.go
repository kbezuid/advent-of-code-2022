package sections

type Sections struct {
	min int
	max int
}

func New(min int, max int) Sections {
	return Sections{
		min: min,
		max: max,
	}
}

func (s *Sections) Contains(min int, max int) bool {
	containsRange := min >= s.min && max <= s.max
	rangeContains := s.min >= min && s.max <= max

	return containsRange || rangeContains
}

func (s *Sections) Overlaps(min int, max int) bool {
	minA := min >= s.min && min <= s.max
	maxA := max >= s.min && max <= s.max

	a := minA || maxA

	minB := s.min >= min && s.min <= max
	maxB := s.max >= min && s.max <= max

	b := minB || maxB

	return a || b
}
