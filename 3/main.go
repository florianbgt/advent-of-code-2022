package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	cli_args := os.Args
	file_path := cli_args[1]
	content := get_file_content(file_path)

	rucksacks := strings.Split(content, "\n")
	sum_priorities := get_priority_amount_from_rucksacks(rucksacks)
	fmt.Println("Sum of priority in each rucksacks is " + strconv.Itoa(sum_priorities))

	groups := get_groups(rucksacks)
	sum_priorities = get_priority_amount_from_groups(groups)
	fmt.Println("Sum of priority in each groups is " + strconv.Itoa(sum_priorities))
}

func get_file_content(file_path string) string {
	file_content, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}
	content := string(file_content)

	return content
}

func get_compartments(items string) (string, string) {
	chunk_size := int(math.Ceil(float64(len(items) / 2)))
	return items[:chunk_size], items[chunk_size:]
}

func get_duplicate_char(list_of_items []string) string {
	number_of_items := len(list_of_items)

	char_count := make(map[string]int)

	for _, items := range list_of_items {
		seen_char := make(map[string]bool)
		for _, item := range items {
			if !seen_char[string(item)] {
				seen_char[string(item)] = true
				char_count[string(item)] += 1
				if char_count[string(item)] == number_of_items {
					return string(item)
				}
			}
		}
	}

	return "!"
}

func get_priority(char string) int {
	chars := "!abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	priority_map := make(map[string]int)

	for i, char := range chars {
		priority_map[string(char)] = i
	}

	return priority_map[char]
}

func get_priority_amount_from_rucksacks(rucksacks []string) int {
	sum_priorities := 0

	for _, rucksack := range rucksacks {
		compartment1, compartment2 := get_compartments(rucksack)

		compartments := []string{compartment1, compartment2}

		duplicate_char := get_duplicate_char(compartments)
		priority := get_priority(duplicate_char)
		sum_priorities += priority
	}

	return sum_priorities
}

func get_groups(items []string) [][]string {
	groups := make([][]string, 0)

	count := 0

	group := make([]string, 0)

	for _, item := range items {
		count += 1
		group = append(group, item)
		if count == 3 {
			groups = append(groups, group)
			group = make([]string, 0)
			count = 0
		}
	}

	return groups
}

func get_priority_amount_from_groups(groups [][]string) int {
	sum_priorities := 0

	for _, group := range groups {
		duplicate_char := get_duplicate_char(group)
		priority := get_priority(duplicate_char)
		sum_priorities += priority
	}

	return sum_priorities
}
