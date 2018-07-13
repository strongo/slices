package slices

import "testing"

func TestRemoveStrings(t *testing.T) {
	t.Parallel()
	s := []string{"v1", "v2", "v3"}

	for _, testCase := range []struct {
		removed int
		source  []string
		remove  []string
		result  []string
	}{
		{removed: 0, source: s, remove: []string{}, result: s},
		{removed: 0, source: s, remove: []string{"v0"}, result: s},
		{removed: 1, source: s, remove: []string{"v1"}, result: []string{"v2", "v3"}},
		{removed: 1, source: s, remove: []string{"v2"}, result: []string{"v1", "v3"}},
		{removed: 1, source: s, remove: []string{"v3"}, result: []string{"v1", "v2"}},
		{removed: 2, source: s, remove: []string{"v1", "v2"}, result: []string{"v3"}},
		{removed: 2, source: s, remove: []string{"v1", "v3"}, result: []string{"v2"}},
		{removed: 2, source: s, remove: []string{"v2", "v3"}, result: []string{"v1"}},
		{removed: 4, source: []string{"v1", "v2", "v2", "v3", "v2", "v2"}, remove: []string{"v2",}, result: []string{"v1", "v3"}},
	} {
		result, removedCount := RemoveStrings(testCase.source, testCase.remove)
		if removedCount != testCase.removed {
			t.Errorf("RemoveStrings(%v, %v) expected to return removedCount=%v, got %v", s, testCase.remove, testCase.removed, removedCount)
		}
		if !EqualStrings(result, testCase.result) {
			t.Errorf("RemoveStrings(%v, %v) expected to return %v, got %v", s, testCase.remove, testCase.result, result)
		}
	}
}

func TestIsInStringSlice(t *testing.T) {
	t.Parallel()
	s := []string{"v1", "v2", "v3"}

	for _, test := range []struct {
		s        []string
		v        string
		expected bool
	}{
		{[]string{}, "v1", false},
		{s, "v0", false},
		{s, "v1", true},
		{s, "v2", true},
		{s, "v3", true},
		{s, "v4", false},
	} {
		if IsInStringSlice(test.v, test.s) != test.expected {
			t.Errorf("Expected %v for IsInStringSlice(%v, %v)", test.expected, test.v, test.s)
		}
	}
}
