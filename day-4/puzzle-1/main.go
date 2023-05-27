package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Advent of Code Puzzle Seventh Puzzle")

	data, error := os.Open("../input.txt")

	check(error)
	scanner := bufio.NewScanner(data)

	sections_group := make([][][]int, 0)
	for scanner.Scan() {
		item := scanner.Text()
		sections := strings.Split(item, ",")

		var sections_int [][]int
		for _, section := range sections {
			item_range := strings.Split(section, "-")
			item_range_int := make([]int, 2)
			for i, item := range item_range {
				item_range_int[i], error = strconv.Atoi(item)
				check(error)
			}
			sections_int = append(sections_int, item_range_int)
		}
		// fmt.Println(sections_int)
		sections_group = append(sections_group, sections_int)
	}

	fully_enclosed_sections_count := 0
	for _, sections := range sections_group {
		first_section, second_section := sections[0], sections[1]

		first_section_size := first_section[1] - first_section[0]
		second_section_size := second_section[1] - second_section[0]

		if first_section_size > second_section_size {
			if first_section[0] <= second_section[0] && first_section[1] >= second_section[1] {
				fmt.Println(sections)
				fmt.Println("First section is bigger")
				fmt.Println("Second section is inside first section")
				fully_enclosed_sections_count++
			}
		} else {
			if second_section[0] <= first_section[0] && second_section[1] >= first_section[1] {
				fmt.Println(sections)
				fmt.Println("Second section is bigger")
				fmt.Println("First section is inside first section")
				fully_enclosed_sections_count++
			}
		}
	}

	fmt.Println("Fully enclosed sections count: ", fully_enclosed_sections_count)
}
