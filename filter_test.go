package slices

import "testing"

func TestRemoveStrings(t *testing.T) {
	t.Parallel()
	s := []string{"v1", "v2", "v3"}

	for _, expects := range []struct {
		changed bool
		remove []string
		result []string
	} {
		{changed: false, remove: []string{}, result: s},
		{changed: false, remove: []string{"v0"}, result: s},
		{changed: true, remove: []string{"v1"}, result: []string{"v2", "v3"}},
		{changed: true, remove: []string{"v2"}, result: []string{"v1", "v3"}},
		{changed: true, remove: []string{"v3"}, result: []string{"v1", "v2"}},
		{changed: true, remove: []string{"v1", "v2"}, result: []string{"v3"}},
		{changed: true, remove: []string{"v1", "v3"}, result: []string{"v2"}},
		{changed: true, remove: []string{"v2", "v3"}, result: []string{"v1"}},
	}{
		result, changed := RemoveStrings(s, expects.remove)
		if changed != expects.changed {
			t.Errorf("RemoveStrings(%v, %v) expected to return changed=%v", s, expects.remove, expects.changed, )
		}
		if !EqualStrings(result, expects.result) {
			t.Errorf("RemoveStrings(%v, %v) expected to return %v, got %v", s, expects.remove, expects.result, result)
		}
	}
}

func TestIsInStringSlice(t *testing.T) {
	t.Parallel()
	s := []string{"v1", "v2", "v3"}

	for _, test := range []struct {
		v        string
		expected bool
	}{
		{"v0", false},
		{"v1", true},
		{"v2", true},
		{"v3", true},
		{"v4", false},
	} {
		if IsInStringSlice(test.v, s) != test.expected {
			t.Errorf("Expected %v for IsInStringSlice(%v, %v)", test.expected, test.v, s)
		}
	}
}
