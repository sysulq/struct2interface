package testdata

// Method describes the code and documentation
// tied into a method
type Method struct {
	Code string
	Docs []string
}

// Lines return a []string consisting of
// the documentation and code appended
// in chronological order
func (m *Method) Lines() []string {
	var lines []string
	lines = append(lines, m.Docs...)
	lines = append(lines, m.Code)
	return lines
}

// Method1 describes the code and documentation
// tied into a method
type Method1 struct {
	Code string
	Docs []string
}

// Lines return a []string consisting of
// the documentation and code appended
// in chronological order
func (m *Method1) Lines() []string {
	var lines []string
	lines = append(lines, m.Docs...)
	lines = append(lines, m.Code)
	return lines
}
