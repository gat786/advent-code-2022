package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func get_points_based_on_choice(choice string) int {
	rock, paper, scissors := "rock", "paper", "scissors"
	rock_points, paper_points, scissors_points := 1, 2, 3

	if choice == rock {
		return rock_points
	} else if choice == paper {
		return paper_points
	} else if choice == scissors {
		return scissors_points
	}
	return 0
}

func calculate_score(choices []string) int {
	win_points, draw_points, lose_points := 6, 3, 0

	rock := []string{"A", "X"}
	paper := []string{"B", "Y"}
	scissors := []string{"C", "Z"}

	opp_choice, my_choice := choices[0], choices[1]

	points_earned := 0
	if slices.Contains(rock, my_choice) {
		if slices.Contains(paper, opp_choice) {
			fmt.Print("Lose")
			points_earned = get_points_based_on_choice("rock") + lose_points
		} else if slices.Contains(scissors, opp_choice) {
			fmt.Print("Win")
			points_earned = get_points_based_on_choice("rock") + win_points
		} else {
			fmt.Print("Draw")
			points_earned = get_points_based_on_choice("rock") + draw_points
		}

	} else if slices.Contains(paper, my_choice) {
		if slices.Contains(rock, opp_choice) {
			fmt.Print("Win")
			points_earned = get_points_based_on_choice("paper") + win_points
		} else if slices.Contains(scissors, opp_choice) {
			fmt.Print("Lose")
			points_earned = get_points_based_on_choice("paper") + lose_points
		} else {
			fmt.Print("Draw")
			points_earned = get_points_based_on_choice("paper") + draw_points
		}
	} else if slices.Contains(scissors, my_choice) {
		if slices.Contains(rock, opp_choice) {
			fmt.Print("Lose")
			points_earned = get_points_based_on_choice("scissors") + lose_points
		} else if slices.Contains(paper, opp_choice) {
			fmt.Print("Win")
			points_earned = get_points_based_on_choice("scissors") + win_points
		} else {
			fmt.Print("Draw")
			points_earned = get_points_based_on_choice("scissors") + draw_points
		}
	}

	fmt.Print("\t Opponent choice: ", opp_choice)
	fmt.Print("\t My choice: ", my_choice)
	fmt.Println("\t Points earned: ", points_earned)

	return points_earned
}

func main() {
	fmt.Println("Advent of Code Puzzle Three")

	data, error := os.Open("../input.txt")

	check(error)

	scanner := bufio.NewScanner(data)
	var total_score int
	for scanner.Scan() {
		line := scanner.Text()

		var choices = strings.Split(line, " ")
		total_score = total_score + calculate_score(choices)
	}

	fmt.Println("Total Score: ", total_score)
}
