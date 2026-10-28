package main

import "testing"

func TestContainsPermutation(t *testing.T) {
	tests := []struct {
		name   string
		s1, s2 string
		want   bool
	}{
		{"example 1", "ab", "eidbaooo", true},
		{"example 2", "ab", "eidboaoo", false},
		{"s1 longer than s2", "abc", "ab", false},
		{"identical strings", "abc", "abc", true},
		{"permutation of whole s2", "abc", "cba", true},
		{"match at start", "abc", "bacxyz", true},
		{"match at end", "abc", "xyzcab", true},
		{"match in middle", "abc", "xxbcaxx", true},
		{"repeated letters match", "aab", "xabaz", true},
		{"repeated letters need exact counts", "aab", "xabxbaxz", false},
		{"single char found", "a", "xyza", true},
		{"single char missing", "a", "xyz", false},
		{"no match same length", "abc", "abd", false},
		{"all same letters", "aaa", "baaab", true},
		{"not contiguous", "ab", "axb", false},
		{"empty s1", "", "abc", true},
		{"both empty", "", "", true},
		{"empty s2", "a", "", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := containsPermutation(tc.s1, tc.s2); got != tc.want {
				t.Errorf("containsPermutation(%q, %q) = %v, want %v", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}
