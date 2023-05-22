package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calculate_score(match_result string, opp_choice string) int {
	rock_choice, paper_choice, scissors_choice := "A", "B", "C"
	win_result, lose_result, draw_result := "Z", "X", "Y"
	rock_point, paper_point, scissors_point := 1, 2, 3
	lose_score, draw_score, win_score := 0, 3, 6
	// fmt.Println("")
	if match_result == win_result {
		fmt.Print("Win")
		if opp_choice == rock_choice {
			fmt.Print("\t Choice made: Paper")
			return paper_point + win_score
		} else if opp_choice == paper_choice {
			fmt.Print("\t Choice made: Scissors")
			return scissors_point + win_score
		} else if opp_choice == scissors_choice {
			fmt.Print("\t Choice made: Rock")
			return rock_point + win_score
		}
	} else if match_result == lose_result {
		fmt.Print("Lose")
		if opp_choice == rock_choice {
			fmt.Print("\t Choice made: Scissors")
			return scissors_point + lose_score
		} else if opp_choice == paper_choice {
			fmt.Print("\t Choice made: rock")
			return rock_point + lose_score
		} else if opp_choice == scissors_choice {
			fmt.Print("\t Choice made: Paper")
			return paper_point + lose_score
		}
	} else if match_result == draw_result {
		fmt.Print("Draw")
		if opp_choice == rock_choice {
			fmt.Print("\t Choice made: Rock")
			return rock_point + draw_score
		} else if opp_choice == paper_choice {
			fmt.Print("\t Choice made: Paper")
			return paper_point + draw_score
		} else if opp_choice == scissors_choice {
			fmt.Print("\t Choice made: Scissors")
			return scissors_point + draw_score
		}
	}
	return 0
}

func main() {
	fmt.Println("Advent of Code Puzzle Fourth")

	data, error := os.Open("../input.txt")

	check(error)

	scanner := bufio.NewScanner(data)

	total_points := 0
	for scanner.Scan() {
		line := scanner.Text()
		choices := strings.Split(line, " ")

		opp_choice, match_result := choices[0], choices[1]
		points_earned := calculate_score(match_result, opp_choice)
		fmt.Println("\t Points earned: ", points_earned)
		total_points = total_points + points_earned
	}
	fmt.Println("Total points: ", total_points)

}
