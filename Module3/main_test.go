package main

import (
	"math/rand"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// normalize sorts each group and then the groups themselves so results can be
// compared regardless of the order groupAnagram returns them in.
func normalize(groups [][]string) [][]string {
	out := make([][]string, len(groups))
	for i, g := range groups {
		out[i] = slices.Clone(g)
		sort.Strings(out[i])
	}
	sort.Slice(out, func(i, j int) bool {
		return strings.Join(out[i], ",") < strings.Join(out[j], ",")
	})
	return out
}

func TestGroupAnagram(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{
			"sample from main",
			[]string{"bucket", "rat", "mango", "tango", "ogtan", "tar"},
			[][]string{{"bucket"}, {"rat", "tar"}, {"mango"}, {"tango", "ogtan"}},
		},
		{"nil input", nil, [][]string{}},
		{"empty input", []string{}, [][]string{}},
		{"single word", []string{"abc"}, [][]string{{"abc"}}},
		{"no anagrams", []string{"abc", "def", "ghi"}, [][]string{{"abc"}, {"def"}, {"ghi"}}},
		{"all anagrams", []string{"listen", "silent", "enlist", "tinsel"}, [][]string{{"enlist", "listen", "silent", "tinsel"}}},
		{"empty strings group together", []string{"", "", "a"}, [][]string{{"", ""}, {"a"}}},
		{"duplicate words kept", []string{"abc", "abc", "bca"}, [][]string{{"abc", "abc", "bca"}}},
		{"repeated letters differ by count", []string{"aab", "abb", "aba"}, [][]string{{"aab", "aba"}, {"abb"}}},
		{"case sensitive", []string{"Ab", "aB", "ba"}, [][]string{{"Ab"}, {"aB"}, {"ba"}}},
		{"different lengths", []string{"a", "aa", "aaa"}, [][]string{{"a"}, {"aa"}, {"aaa"}}},
		{"digits and punctuation", []string{"a1!", "!1a", "1a"}, [][]string{{"!1a", "a1!"}, {"1a"}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := normalize(groupAnagram(tc.input))
			want := normalize(tc.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("groupAnagram(%q) = %q, want %q", tc.input, got, want)
			}
		})
	}
}

func TestGroupAnagramPreservesAllInputs(t *testing.T) {
	input := []string{"bucket", "rat", "mango", "tango", "ogtan", "tar", "rat"}
	var flat []string
	for _, g := range groupAnagram(input) {
		flat = append(flat, g...)
	}
	sort.Strings(flat)
	want := slices.Clone(input)
	sort.Strings(want)
	if !reflect.DeepEqual(flat, want) {
		t.Errorf("words lost or invented: got %q, want %q", flat, want)
	}
}

func TestSortString(t *testing.T) {
	tests := []struct{ in, want string }{
		{"", ""},
		{"a", "a"},
		{"mango", "agmno"},
		{"ogtan", "agnot"},
		{"bucket", "bcektu"},
		{"aaa", "aaa"},
		{"banana", "aaabnn"},
		{"dcba", "abcd"},
		{"abcd", "abcd"},
		{"bA", "Ab"}, // uppercase sorts before lowercase in ASCII
		{"a1!", "!1a"},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			if got := sortString(tc.in); got != tc.want {
				t.Errorf("sortString(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

// sorters lists the algorithm implementations under test.
var sorters = []struct {
	name string
	fn   func([]rune) []rune
}{
	{"quicksort", quicksort},
	{"radixSort", radixSort},
}

func runesSorted(r []rune) []rune {
	out := slices.Clone(r)
	slices.Sort(out)
	return out
}

func TestSortingAlgorithms(t *testing.T) {
	cases := []struct {
		name string
		in   []rune
	}{
		{"empty", []rune{}},
		{"single", []rune("a")},
		{"two sorted", []rune("ab")},
		{"two reversed", []rune("ba")},
		{"already sorted", []rune("abcdef")},
		{"reverse sorted", []rune("fedcba")},
		{"all equal", []rune("aaaaa")},
		{"duplicates", []rune("banana")},
		{"mixed case and symbols", []rune("Hello, World! 123")},
		{"full printable ASCII reversed", func() []rune {
			var r []rune
			for c := '~'; c >= ' '; c-- {
				r = append(r, c)
			}
			return r
		}()},
	}

	for _, s := range sorters {
		for _, tc := range cases {
			t.Run(s.name+"/"+tc.name, func(t *testing.T) {
				got := s.fn(slices.Clone(tc.in))
				want := runesSorted(tc.in)
				if !slices.Equal(got, want) {
					t.Errorf("%s(%q) = %q, want %q", s.name, string(tc.in), string(got), string(want))
				}
			})
		}
	}
}

func TestSortingAlgorithmsAgainstReference(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for _, s := range sorters {
		t.Run(s.name, func(t *testing.T) {
			for range 500 {
				in := make([]rune, rng.Intn(60))
				for i := range in {
					in[i] = rune('a' + rng.Intn(6)) // small alphabet forces many duplicates
				}
				got := s.fn(slices.Clone(in))
				want := runesSorted(in)
				if !slices.Equal(got, want) {
					t.Fatalf("%s(%q) = %q, want %q", s.name, string(in), string(got), string(want))
				}
			}
		})
	}
}

func TestSortingDoesNotMutateInput(t *testing.T) {
	for _, s := range sorters {
		t.Run(s.name, func(t *testing.T) {
			in := []rune("dcba")
			s.fn(in)
			if string(in) != "dcba" {
				t.Errorf("%s mutated its input: %q", s.name, string(in))
			}
		})
	}
}

func TestGetMax(t *testing.T) {
	if got := getMax([]rune("abcz")); got != int('z') {
		t.Errorf("getMax = %d, want %d", got, 'z')
	}
	if got := getMax([]rune("a")); got != int('a') {
		t.Errorf("getMax single = %d, want %d", got, 'a')
	}
}

func TestCountingSort(t *testing.T) {
	// 'a'=97, 'b'=98, 'c'=99, so exp=1 orders by the ones digit.
	got := countingSort([]rune("cab"), 1)
	if want := []rune("abc"); !slices.Equal(got, want) {
		t.Errorf("countingSort exp=1 = %q, want %q", string(got), string(want))
	}
	// 'a' and 'b' share a tens digit (9), so a stable sort keeps input order.
	got = countingSort([]rune("ba"), 10)
	if want := []rune("ba"); !slices.Equal(got, want) {
		t.Errorf("countingSort not stable: got %q, want %q", string(got), string(want))
	}
}

// Benchmarks compare the two sorting algorithms on identical inputs.
// Run with: go test -bench=. -benchmem
func BenchmarkSorters(b *testing.B) {
	for _, n := range []int{10, 100, 1000, 10000} {
		rng := rand.New(rand.NewSource(42))
		in := make([]rune, n)
		for i := range in {
			in[i] = rune('a' + rng.Intn(26))
		}
		for _, s := range sorters {
			b.Run(s.name+"/n="+strconv.Itoa(n), func(b *testing.B) {
				for b.Loop() {
					s.fn(slices.Clone(in))
				}
			})
		}
	}
}
