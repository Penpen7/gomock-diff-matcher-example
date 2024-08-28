package matcher

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	gomock "go.uber.org/mock/gomock"
)

// diffEq is a custom matcher that uses go-cmp for comparison
type diffEq struct {
	matcher  gomock.Matcher
	expected interface{}
	actual   interface{}
}

// Matches implements the gomock.Matcher interface
func (m *diffEq) Matches(x interface{}) bool {
	m.actual = x
	return m.matcher.Matches(x)
}

// String implements the gomock.Matcher interface
func (m *diffEq) String() string {
	if m.actual == nil {
		return fmt.Sprintf("is equal to %v", m.expected)
	}
	return fmt.Sprintf("is equal to %v, but got %v\nDiff: %s(-want, +got)", m.expected, m.actual, cmp.Diff(m.expected, m.actual))
}

// DiffEq creates a new instance of CustomMatcher
func DiffEq(expected interface{}) *diffEq {
	return &diffEq{matcher: gomock.Eq(expected), expected: expected}
}
