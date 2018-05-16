package slices

import "testing"

func TestEqualStrings(t *testing.T) {
	for _, test := range []struct {
		expected bool
		v1       []string
		v2       []string
	}{
		{expected: true, v1: []string{}, v2: []string{}},
		{expected: false, v1: []string{"v1"}, v2: []string{}},
		{expected: false, v1: []string{}, v2: []string{"v2"}},
		{expected: false, v1: []string{"v1"}, v2: []string{"v2"}},
		{expected: true, v1: []string{"v1"}, v2: []string{"v1"}},
		{expected: true, v1: []string{"v1", "v2"}, v2: []string{"v1", "v2"}},
		{expected: false, v1: []string{"v1", "v2"}, v2: []string{"v2", "v1"}},
		{expected: true, v1: []string{"v1", "v2", "v3"}, v2: []string{"v1", "v2", "v3"}},
	} {
		if test.expected != EqualStrings(test.v1, test.v2) {
			t.Errorf("Expected %v for EqualStrings(%v, %v)", test.expected, test.v1, test.v2)
		}
	}
}
