package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Advent of Code Puzzle One")

	data, error := os.Open("../input.txt")

	check(error)

	scanner := bufio.NewScanner(data)

	var number_of_elves int
	var current_elf_total int

	var elves_total_calories []int

	for scanner.Scan() {
		if scanner.Text() == "" {
			elves_total_calories = append(elves_total_calories, current_elf_total)
			current_elf_total = 0
			number_of_elves++
		} else {
			calorie, err := strconv.Atoi(scanner.Text())
			check(err)
			current_elf_total = current_elf_total + calorie
		}
	}

	sort.Ints(elves_total_calories)

	fmt.Println("Sorted calories: ", elves_total_calories)
	var last_three []int = elves_total_calories[len(elves_total_calories)-3:]
	fmt.Println("Last three: ", last_three)

	var sum int
	for i := 0; i < len(last_three); i++ {
		sum = sum + last_three[i]
	}

	fmt.Println("Sum of last three: ", sum)
}
