package main

import (
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"testing"
)

// buildGraph creates cities from a population map and connects them by edges.
func buildGraph(pops map[string]int, edges [][2]string) map[string]*City {
	m := make(map[string]*City)
	for name, pop := range pops {
		m[name] = &City{Name: name, Population: pop, Connections: []*City{}}
	}
	for _, e := range edges {
		m[e[0]].AddConnection(m[e[1]])
	}
	return m
}

func writeFile(t *testing.T, name, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

func hasConnection(c *City, name string) bool {
	for _, n := range c.Connections {
		if n.Name == name {
			return true
		}
	}
	return false
}

// Two islands: {A,B,C,D} (A-B, B-C, C-D) and {E,F} (E-F), plus isolated-by-edge island {G}.
func sampleGraph() map[string]*City {
	return buildGraph(
		map[string]int{"A": 100, "B": 200, "C": 300, "D": 400, "E": 50, "F": 60, "G": 7},
		[][2]string{{"A", "B"}, {"B", "C"}, {"C", "D"}, {"E", "F"}},
	)
}

// ---- Task 1: City ----

func TestCityFields(t *testing.T) {
	c := &City{Name: "X", Population: 42, Connections: []*City{}}
	if c.Name != "X" || c.Population != 42 || len(c.Connections) != 0 {
		t.Fatalf("unexpected city: %+v", c)
	}
}

func TestAddConnectionIsBidirectional(t *testing.T) {
	a := &City{Name: "A"}
	b := &City{Name: "B"}
	a.AddConnection(b)
	if !hasConnection(a, "B") {
		t.Error("A should be connected to B")
	}
	if !hasConnection(b, "A") {
		t.Error("B should be connected to A (highways are bi-directional)")
	}
}

// ---- Task 2: reading files / building the graph ----

func TestReadCitiesFromFile(t *testing.T) {
	tests := []struct {
		name    string
		content string
	}{
		{"spec format", "City A: 1000\nCity B: 2000\n"},
		{"spaced colon (as in city_population.txt)", "City A : 1000\nCity B : 2000\n"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			path := writeFile(t, "pop.txt", tc.content)
			m, err := readCitiesFromFile(path)
			if err != nil {
				t.Fatal(err)
			}
			if len(m) != 2 {
				t.Fatalf("expected 2 cities, got %d: %v", len(m), m)
			}
			if c := m["City A"]; c == nil || c.Population != 1000 {
				t.Errorf("City A wrong: %+v", c)
			}
			if c := m["City B"]; c == nil || c.Population != 2000 {
				t.Errorf("City B wrong: %+v", c)
			}
		})
	}
}

func TestReadCitiesFromFileSkipsMalformedLines(t *testing.T) {
	path := writeFile(t, "pop.txt", "Good: 10\nno separator\nBad: notanumber\nAlso Good: 20\n")
	m, err := readCitiesFromFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(m) != 2 || m["Good"] == nil || m["Also Good"] == nil {
		t.Errorf("expected only valid cities, got %v", m)
	}
}

func TestReadCitiesFromMissingFile(t *testing.T) {
	if _, err := readCitiesFromFile(filepath.Join(t.TempDir(), "nope.txt")); err == nil {
		t.Error("expected error for missing file")
	}
}

func TestReadConnectionsFromFile(t *testing.T) {
	for _, sep := range []string{": ", " : "} {
		t.Run("sep"+sep, func(t *testing.T) {
			m := buildGraph(map[string]int{"A": 1, "B": 2, "C": 3}, nil)
			path := writeFile(t, "roads.txt", "A"+sep+"B\nB"+sep+"C\n")
			if err := readConnectionsFromFile(path, m); err != nil {
				t.Fatal(err)
			}
			if !hasConnection(m["A"], "B") || !hasConnection(m["B"], "A") {
				t.Error("A<->B missing")
			}
			if !hasConnection(m["B"], "C") || !hasConnection(m["C"], "B") {
				t.Error("B<->C missing")
			}
			if hasConnection(m["A"], "C") {
				t.Error("A and C should not be directly connected")
			}
		})
	}
}

func TestReadConnectionsIgnoresUnknownCities(t *testing.T) {
	m := buildGraph(map[string]int{"A": 1, "B": 2}, nil)
	path := writeFile(t, "roads.txt", "A: Ghost\nA: B\n")
	if err := readConnectionsFromFile(path, m); err != nil {
		t.Fatal(err)
	}
	if len(m["A"].Connections) != 1 || !hasConnection(m["A"], "B") {
		t.Errorf("A connections wrong: %v", m["A"].Connections)
	}
}

func TestReadConnectionsFromMissingFile(t *testing.T) {
	m := buildGraph(map[string]int{"A": 1}, nil)
	if err := readConnectionsFromFile(filepath.Join(t.TempDir(), "nope.txt"), m); err == nil {
		t.Error("expected error for missing file")
	}
}

func TestReadActualDataFiles(t *testing.T) {
	m, err := readCitiesFromFile("city_population.txt")
	if err != nil {
		t.Fatal(err)
	}
	if len(m) == 0 {
		t.Fatal("no cities parsed from city_population.txt")
	}
	if c := m["New York"]; c == nil || c.Population != 8405837 {
		t.Errorf("New York wrong: %+v", c)
	}
	if err := readConnectionsFromFile("road_network.txt", m); err != nil {
		t.Fatal(err)
	}
	if !hasConnection(m["Citrus Heights"], "Rialto") {
		t.Error("Citrus Heights should connect to Rialto")
	}
}

// ---- Task 3: number of islands ----

func TestNumberOfIslands(t *testing.T) {
	islands := numberOfIslands(sampleGraph())
	if len(islands) != 3 {
		t.Fatalf("expected 3 islands, got %d", len(islands))
	}
}

func TestNumberOfIslandsEmpty(t *testing.T) {
	if n := len(numberOfIslands(map[string]*City{})); n != 0 {
		t.Errorf("expected 0 islands, got %d", n)
	}
}

func TestNumberOfIslandsSingleConnected(t *testing.T) {
	m := buildGraph(map[string]int{"A": 1, "B": 2, "C": 3}, [][2]string{{"A", "B"}, {"B", "C"}})
	if n := len(numberOfIslands(m)); n != 1 {
		t.Errorf("expected 1 island, got %d", n)
	}
}

func TestNumberOfIslandsMultiEdgesAndCycle(t *testing.T) {
	m := buildGraph(map[string]int{"A": 1, "B": 2, "C": 3},
		[][2]string{{"A", "B"}, {"A", "B"}, {"B", "C"}, {"C", "A"}})
	if n := len(numberOfIslands(m)); n != 1 {
		t.Errorf("expected 1 island, got %d", n)
	}
}

func TestIslandMembership(t *testing.T) {
	islands := numberOfIslands(sampleGraph())
	var sizes []int
	for _, is := range islands {
		sizes = append(sizes, len(is))
	}
	sort.Ints(sizes)
	want := []int{1, 2, 4}
	if len(sizes) != len(want) {
		t.Fatalf("island sizes %v, want %v", sizes, want)
	}
	for i := range want {
		if sizes[i] != want[i] {
			t.Errorf("island sizes %v, want %v", sizes, want)
			break
		}
	}
}

// ---- Task 4: island populations ----

func TestTotalPopulation(t *testing.T) {
	if got := totalPopulation(nil); got != 0 {
		t.Errorf("empty island = %d, want 0", got)
	}
	island := []*City{{Population: 10}, {Population: 20}, {Population: 30}}
	if got := totalPopulation(island); got != 60 {
		t.Errorf("got %d, want 60", got)
	}
}

func TestIslandPopulations(t *testing.T) {
	islands := numberOfIslands(sampleGraph())
	var pops []int
	for _, is := range islands {
		pops = append(pops, totalPopulation(is))
	}
	sort.Ints(pops)
	want := []int{7, 110, 1000}
	if len(pops) != len(want) {
		t.Fatalf("populations %v, want %v", pops, want)
	}
	for i := range want {
		if pops[i] != want[i] {
			t.Errorf("populations %v, want %v", pops, want)
			break
		}
	}
}

// ---- Task 5: distance between cities ----

func TestDistanceTo(t *testing.T) {
	m := sampleGraph()
	tests := []struct {
		from, to string
		want     int
	}{
		{"A", "A", 0},
		{"A", "B", 1},
		{"A", "C", 2},
		{"A", "D", 3},
		{"D", "A", 3}, // bi-directional
		{"A", "E", -1},
		{"G", "A", -1},
	}
	for _, tc := range tests {
		if got := m[tc.from].DistanceTo(m[tc.to]); got != tc.want {
			t.Errorf("%s -> %s = %d, want %d", tc.from, tc.to, got, tc.want)
		}
	}
}

func TestDistanceToPicksShortestPath(t *testing.T) {
	// A-B-C-D-E long path, plus shortcut A-E.
	m := buildGraph(map[string]int{"A": 1, "B": 1, "C": 1, "D": 1, "E": 1},
		[][2]string{{"A", "B"}, {"B", "C"}, {"C", "D"}, {"D", "E"}, {"A", "E"}})
	if got := m["A"].DistanceTo(m["E"]); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
	if got := m["B"].DistanceTo(m["E"]); got != 2 {
		t.Errorf("got %d, want 2", got)
	}
}

func TestDistanceToWithDuplicateRoadsAndCycles(t *testing.T) {
	m := buildGraph(map[string]int{"A": 1, "B": 1, "C": 1},
		[][2]string{{"A", "B"}, {"A", "B"}, {"B", "C"}, {"C", "A"}})
	if got := m["A"].DistanceTo(m["C"]); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
}

// diamondChain builds `diamonds` stacked diamonds: top -> (l,r) -> next top.
// The first city is "t0", the last is "t<diamonds>"; the distance between them is 2*diamonds
// and there are 2^diamonds shortest paths.
func diamondChain(diamonds int) (*City, *City) {
	pops := map[string]int{}
	var edges [][2]string
	name := func(prefix string, i int) string { return prefix + strconv.Itoa(i) }
	for i := 0; i <= diamonds; i++ {
		pops[name("t", i)] = 1
	}
	for i := 0; i < diamonds; i++ {
		pops[name("l", i)] = 1
		pops[name("r", i)] = 1
		edges = append(edges,
			[2]string{name("t", i), name("l", i)}, [2]string{name("t", i), name("r", i)},
			[2]string{name("l", i), name("t", i+1)}, [2]string{name("r", i), name("t", i+1)})
	}
	m := buildGraph(pops, edges)
	return m["t0"], m[name("t", diamonds)]
}

func TestDistanceToManyShortestPaths(t *testing.T) {
	const diamonds = 27
	from, to := diamondChain(diamonds)
	if got := from.DistanceTo(to); got != 2*diamonds {
		t.Errorf("got %d, want %d", got, 2*diamonds)
	}
}
