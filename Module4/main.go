package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type City struct {
	Name        string
	Population  int
	Connections []*City
}

func (c *City) AddConnection(other *City) {
	c.Connections = append(c.Connections, other)
	other.Connections = append(other.Connections, c)
}

func (c *City) DistanceTo(other *City) int {
	// bfs to find the shortest path
	visited := make(map[string]bool)
	queue := []*City{c}
	visited[c.Name] = true
	distance := 0

	for len(queue) > 0 {
		nextQueue := []*City{}
		for _, city := range queue {
			if city.Name == other.Name {
				return distance
			}
			for _, neighbor := range city.Connections {
				if !visited[neighbor.Name] {
					nextQueue = append(nextQueue, neighbor)
					visited[neighbor.Name] = true
				}
			}
		}
		queue = nextQueue
		distance++
	}
	return -1
}

func main() {
	cityMap, err := readCitiesFromFile("city_population.txt")
	if err != nil {
		fmt.Println("Error reading cities from file:", err)
		return
	}

	err = readConnectionsFromFile("road_network.txt", cityMap)
	if err != nil {
		fmt.Println("Error reading connections from file:", err)
		return
	}

	for _, city := range cityMap {
		fmt.Printf("City: %s, Population: %d, Connections: ", city.Name, city.Population)
		for _, conn := range city.Connections {
			fmt.Printf("%s ", conn.Name)
		}
		fmt.Println()
	}

	// Count the number of islands
	islands := numberOfIslands(cityMap)
	fmt.Printf("Number of islands: %d\n", len(islands))
	for i, island := range islands {
		fmt.Printf("Total Population of Island %d: %d\n", i+1, totalPopulation(island))
	}
}

func readCitiesFromFile(filename string) (map[string]*City, error) {
	nameMap := make(map[string]*City)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}
		name := strings.TrimSpace(parts[0])
		population, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			fmt.Println("Invalid population value:", parts[1])
			continue
		}
		nameMap[name] = &City{Name: name, Population: population, Connections: []*City{}}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return nameMap, nil
}

func readConnectionsFromFile(filename string, cityMap map[string]*City) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}
		city1Name := strings.TrimSpace(parts[0])
		city2Name := strings.TrimSpace(parts[1])

		city1, ok1 := cityMap[city1Name]
		city2, ok2 := cityMap[city2Name]
		if !ok1 || !ok2 {
			fmt.Println("One or both cities not found:", city1Name, city2Name)
			continue
		}
		city1.AddConnection(city2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
	return nil
}

func numberOfIslands(cityMap map[string]*City) [][]*City {
	visited := make(map[string]struct{})
	res := [][]*City{}

	var dfs func(city *City, visited map[string]bool)

	dfs = func(city *City, visited map[string]bool) {
		visited[city.Name] = true
		for _, neighbor := range city.Connections {
			if !visited[neighbor.Name] {
				dfs(neighbor, visited)
			}
		}
	}

	for _, city := range cityMap {
		if _, ok := visited[city.Name]; !ok {
			island := []*City{}
			islandCityVisited := make(map[string]bool)
			dfs(city, islandCityVisited)
			for name := range islandCityVisited {
				island = append(island, cityMap[name])
				visited[name] = struct{}{}
			}
			res = append(res, island)
		}
	}

	return res
}

func totalPopulation(island []*City) int {
	total := 0
	for _, city := range island {
		total += city.Population
	}
	return total
}
