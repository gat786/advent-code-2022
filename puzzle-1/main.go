package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Advent of Code Puzzle One")

	data, error := os.Open("input.txt")

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

	fmt.Println("Number of elves: ", number_of_elves)

	var elf_with_max_calories_index int
	var max_calories int
	for i := 0; i < len(elves_total_calories); i++ {
		if elves_total_calories[i] > max_calories {
			max_calories = elves_total_calories[i]
			elf_with_max_calories_index = i
		}
		// fmt.Println("Elf ", i, " has ", elves_total_calories[i], " calories.")
	}

	fmt.Println("Elf number", elf_with_max_calories_index, "has", max_calories, "calories.")
}
