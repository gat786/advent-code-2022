package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func get_priorities() map[string]int {
	chars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i",
		"j", "k", "l", "m", "n", "o", "p", "q", "r",
		"s", "t", "u", "v", "w", "x", "y", "z", "A",
		"B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S",
		"T", "U", "V", "W", "X", "Y", "Z"}

	priorities := make(map[string]int)

	for i := 0; i < len(chars); i++ {
		priorities[chars[i]] = i + 1
	}

	return priorities
}

func main() {
	fmt.Println("Advent of Code Puzzle Fifth")

	data, error := os.Open("../input.txt")

	check(error)

	priorites := get_priorities()
	fmt.Println(priorites)

	scanner := bufio.NewScanner(data)
	var priorites_present []int

	for scanner.Scan() {
		line := scanner.Text()

		no_of_items := len(line)
		fmt.Println("count of items in this rucksack are: ", no_of_items)
		items_in_each_comp := no_of_items / 2

		upper_comp := line[:items_in_each_comp]
		lower_comp := line[items_in_each_comp:]

		fmt.Println("upper comp: ", upper_comp, "lower comp: ", lower_comp)

		items_in_first_comp := make(map[string]int)

		for i := 0; i < len(upper_comp); i++ {
			charachter := string(upper_comp[i])
			items_in_first_comp[charachter] = priorites[charachter]
		}

		for i := 0; i < len(lower_comp); i++ {
			item := string(lower_comp[i])
			if _, ok := items_in_first_comp[item]; ok {
				priority := priorites[item]
				fmt.Println("item: ", item, "is in both which has a priority: ", priority)
				priorites_present = append(priorites_present, priorites[item])
				break
			}
		}
	}
	fmt.Println("priorities present: ", priorites_present)

	total_priority := 0
	for i := 0; i < len(priorites_present); i++ {
		total_priority += priorites_present[i]
	}

	fmt.Println("total priority: ", total_priority)
	// return total_priority
}
