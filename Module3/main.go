package main

import (
	"fmt"
)

func main() {
	// Example usage
	strings := []string{"bucket", "rat", "mango", "tango", "ogtan", "tar"}
	anagramGroups := groupAnagram(strings)
	fmt.Println(anagramGroups)
}

func groupAnagram(strings []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, str := range strings {
		sortedStr := sortString(str)
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}

	var result [][]string
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func sortString(str string) string {
	runes := []rune(str)

	// Use quicksort to sort the runes
	// runes = quicksort(runes)

	// Use radix sort to sort the runes
	runes = radixSort(runes)
	return string(runes)
}

func quicksort(runes []rune) []rune {
	if len(runes) < 2 {
		return runes
	}

	pivot := runes[len(runes)/2]
	left := []rune{}
	middle := []rune{}
	right := []rune{}

	for _, r := range runes {
		if r < pivot {
			left = append(left, r)
		} else if r > pivot {
			right = append(right, r)
		} else {
			middle = append(middle, r)
		}
	}

	return append(append(quicksort(left), middle...), quicksort(right)...)
}

func radixSort(runes []rune) []rune {
	if len(runes) == 0 {
		return runes
	}
	max := getMax(runes)

	for exp := 1; max/exp > 0; exp *= 10 {
		runes = countingSort(runes, exp)
	}
	return runes
}

func getMax(runes []rune) int {
	max := int(runes[0])
	for _, r := range runes {
		if int(r) > max {
			max = int(r)
		}
	}
	return max
}

func countingSort(runes []rune, exp int) []rune {
	n := len(runes)
	output := make([]rune, n)
	count := make([]int, 10)

	for i := range n {
		count[(int(runes[i])/exp)%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		index := (int(runes[i]) / exp) % 10
		output[count[index]-1] = runes[i]
		count[index]--
	}
	return output
}
