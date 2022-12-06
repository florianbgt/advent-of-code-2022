package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cli_args := os.Args
	file_path := cli_args[1]
	content := get_file_content(file_path)

	content_list := strings.Split(content, "\n\n")

	initial_position := content_list[0]
	list_initial_position := strings.Split(initial_position, "\n")

	stacks := get_stacks(list_initial_position)

	instructions_content := content_list[1]
	list_instructions_string := strings.Split(instructions_content, "\n")

	instructions := get_instrutions(list_instructions_string)

	executed_stacks := execute_instruction(instructions, stacks)

	last_chars_stacks := ""
	for _, stack := range executed_stacks {
		last_chars_stacks += stack[len(stack)-1]
	}

	fmt.Println("9000 final state: " + last_chars_stacks)

	stacks = get_stacks(list_initial_position)

	executed_stacks_9001 := execute_instruction_9001(instructions, stacks)

	last_chars_stacks2 := ""
	for _, stack := range executed_stacks_9001 {
		last_chars_stacks2 += stack[len(stack)-1]
	}

	fmt.Println("9001 final state: " + last_chars_stacks2)
}

func get_file_content(file_path string) string {
	file_content, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}
	content := string(file_content)

	return content
}

func get_stacks(list_initial_position []string) [][]string {
	stacks := make([][]string, 0)

	for i := 0; i < len(list_initial_position)-1; i++ {
		line := list_initial_position[i]
		counter_col := 0
		last := len(line)

		for j := 0; j <= last; {
			chars := strings.Split(line, "")
			char := chars[j+1]
			if i == 0 {
				stacks = append(stacks, make([]string, 0))
			}
			str_char := string(char)
			if str_char != " " {
				stacks[counter_col] = append([]string{str_char}, stacks[counter_col]...)
			}
			j += 4
			counter_col += 1
		}
	}

	return stacks
}

func get_instrutions(list_instructions_string []string) []map[string]int {
	list_instructions_dict := make([]map[string]int, 0)

	for _, instruction_string := range list_instructions_string {
		instruction_list := strings.Split(instruction_string, " ")
		amount, err := strconv.Atoi(string(instruction_list[1]))
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(string(instruction_list[3]))
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(string(instruction_list[5]))
		if err != nil {
			panic(err)
		}
		list_instructions_dict = append(list_instructions_dict, map[string]int{"amount": amount, "from": from, "to": to})
	}

	return list_instructions_dict
}

func execute_instruction(instructions []map[string]int, stacks [][]string) [][]string {
	for _, instruction := range instructions {
		amount := instruction["amount"] - 1
		from := instruction["from"] - 1
		to := instruction["to"] - 1

		for i := 0; i <= amount; i++ {
			taken := stacks[from][len(stacks[from])-1]
			stacks[to] = append(stacks[to], taken)
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}

	return stacks
}

func execute_instruction_9001(instructions []map[string]int, stacks [][]string) [][]string {
	for _, instruction := range instructions {
		amount := instruction["amount"]
		from := instruction["from"] - 1
		to := instruction["to"] - 1

		taken := stacks[from][len(stacks[from])-amount:]

		stacks[to] = append(stacks[to], taken...)
		stacks[from] = stacks[from][:len(stacks[from])-amount]
	}

	return stacks
}
