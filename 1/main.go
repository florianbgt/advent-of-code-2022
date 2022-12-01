package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content := get_file_content("/1/input.txt")

	elfs := strings.Split(content, "\n\n")

	number_of_elfs := 3

	sum_calories := get_total_calories_for_x_efs(elfs, number_of_elfs)

	fmt.Println("\nThe sum of the calories of the top " + strconv.Itoa(number_of_elfs) + " elves with the most calories is " + strconv.Itoa(sum_calories) + ".")
}

func get_file_content(file_path string) string {
	directory, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file_content, err := os.ReadFile(directory + file_path)
	if err != nil {
		panic(err)
	}
	content := string(file_content)

	return content
}

func get_max_calories_from_elfs(elfs []string) (int, int) {
	max_calories := 0
	max_calories_index := 0

	for index, elf := range elfs {
		items := strings.Split(elf, "\n")
		total_calories := 0
		for _, item := range items {
			calory_amount, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			total_calories += calory_amount
		}
		if total_calories > max_calories {
			max_calories = total_calories
			max_calories_index = index
		}
	}

	return max_calories, max_calories_index
}

func remove_elf(elfs []string, index int) []string {
	elfs = append(elfs[:index], elfs[index+1:]...)
	return elfs
}

func get_total_calories_for_x_efs(elfs []string, x int) int {
	sum_calories := 0

	for i := 1; i <= x; i++ {
		max_calories, max_calories_index := get_max_calories_from_elfs(elfs)

		sum_calories += max_calories

		fmt.Println("The elf with the " + strconv.Itoa(i) + " most calories has " + strconv.Itoa(max_calories) + " calories.")

		elfs = remove_elf(elfs, max_calories_index)
	}

	return sum_calories
}
