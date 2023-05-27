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

func update_occurrence(map_to_update map[string]int, item string) map[string]int {
	val, ok := map_to_update[item]
	if ok {
		map_to_update[item] = val + 1
	} else {
		map_to_update[item] = 1
	}

	return map_to_update
}

func main() {
	fmt.Println("Advent of Code Puzzle Sixth Puzzle")

	data, error := os.Open("../input.txt")

	check(error)
	scanner := bufio.NewScanner(data)

	var groups [][]string
	var last_item []string
	for scanner.Scan() {
		item := scanner.Text()

		if len(last_item) < 3 {
			last_item = append(last_item, item)
		} else {
			groups = append(groups, last_item)
			last_item = []string{}
			last_item = append(last_item, item)
		}
	}

	groups = append(groups, last_item)

	var common_items []string
	for _, group := range groups {
		group_1 := group[0]
		group_2 := group[1]
		group_3 := group[2]

		fmt.Println("group: ", group_1, group_2, group_3)

		// fmt.Println("group: ", len(group))
		initial_map := make(map[string]int)
		for _, char := range group_1 {
			initial_map[string(char)] = 1
		}

		second_map := make(map[string]int)
		for _, char := range group_2 {
			_, ok := initial_map[string(char)]
			if ok {
				second_map[string(char)] = 1
			}
		}

		fmt.Println("common items after checking second elf: ", second_map)

		third_list := []string{}
		for _, char := range group_3 {
			_, ok := second_map[string(char)]
			if ok {
				third_list = append(third_list, string(char))
				break
			}
		}

		fmt.Println("common items after checking third elf: ", third_list)

		// check if third map is empty or not
		if len(third_list) == 0 {
			fmt.Println("third list is empty")
		}

		common_items = append(common_items, third_list[0])

	}

	fmt.Println(common_items)

	priorites := get_priorities()
	sum_of_priorities := 0
	for _, item := range common_items {
		fmt.Println("item: ", item)
		p := priorites[item]
		fmt.Println("item: ", item, "priority: ", p)
		sum_of_priorities += p
		fmt.Println("current total priorities: ", sum_of_priorities)
	}

	// fmt.Println("sum of priorities: ", sum_of_priorities)
	// occurrences := make(map[string]int)
}
