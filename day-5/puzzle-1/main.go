package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	values []rune
}

func (s *Stack) Push(value rune) {
	s.values = append(s.values, value)
}

func (s *Stack) Pop() rune {
	value := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return value
}

func (s *Stack) Peek() rune {
	return s.values[len(s.values)-1]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ChunkString(s string, chunkSize int) []string {
	var chunks []string
	runes := []rune(s)

	if len(runes) == 0 {
		return []string{s}
	}

	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]))
	}
	return chunks
}

func perform_move(item_stacks []Stack, move_from int, move_to int, item_count int) {
	stack_to_move_from := item_stacks[move_from]
	stack_to_move_to := item_stacks[move_to]
	for i := 0; i < item_count; i++ {
		item := stack_to_move_from.Pop()
		stack_to_move_to.Push(item)
	}
	item_stacks[move_from] = stack_to_move_from
	item_stacks[move_to] = stack_to_move_to
}

func print_stacks(item_stacks []Stack) {
	for index, item_stack := range item_stacks {
		fmt.Print(index+1, ": ")
		for _, item := range item_stack.values {
			fmt.Print(string(item))
		}
		fmt.Println()
	}
}

func print_top_crate(item_stacks []Stack) {
	for index, item_stack := range item_stacks {
		fmt.Print(index+1, ": ")
		fmt.Print(string(item_stack.Peek()))
		fmt.Println()
	}
}

func main() {
	fmt.Println("Advent of Code Puzzle Ninth Puzzle")

	data, error := os.Open("../input.txt")

	check(error)

	scanner := bufio.NewScanner(data)

	moves := make([]string, 0)
	// stacks := make([]stack, 0)
	stack_info := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "move") {
			moves = append(moves, text)
		} else {
			stack_info = append(stack_info, text)
		}
	}
	total_lines := len(stack_info) - 2
	// total_chunks_in_a_line := len(ChunkString(stack_info[0], 4))
	stacks_list := [][]string{}
	for line_number, line := range stack_info[:len(stack_info)-2] {
		stacks_list = append(stacks_list, []string{})
		chunks := ChunkString(line, 4)
		for _, chunk := range chunks {
			item := strings.TrimSpace(chunk)
			// var item_char string
			if len(item) > 1 {
				// fmt.Println("line number", line_number, "has", string(item[1]), "at chunk number", chunk_number)
				stacks_list[line_number] = append(stacks_list[line_number], string(item[1]))
			} else {
				stacks_list[line_number] = append(stacks_list[line_number], "_")
			}
		}
	}

	stacks := make([]Stack, total_lines+1)
	// fmt.Println(stacks)
	for i := len(stacks_list) - 1; i >= 0; i-- {
		line := stacks_list[i]
		for stack_number, item := range line {
			if item == "_" {
				continue
			}
			stacks[stack_number].Push(rune(item[0]))
		}
		// fmt.Println()
	}

	fmt.Println("original stack :")
	print_stacks(stacks)
	for _, line := range moves {
		fmt.Println(line)
		chunks := strings.Split(line, " ")
		move_items, err := strconv.Atoi(chunks[1])
		check(err)
		move_from, err := strconv.Atoi(chunks[3])
		check(err)
		move_to, err := strconv.Atoi(chunks[5])
		check(err)
		// fmt.Println(move_items, move_from, move_to)
		perform_move(stacks, move_from-1, move_to-1, move_items)
		print_stacks(stacks)
	}
	// fmt.Println(stacks)
	fmt.Println("top crates would be :")
	print_top_crate(stacks)
}
