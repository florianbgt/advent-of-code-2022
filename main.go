package main

import (
	"fmt"
	"os"
	"strconv"

	solvers "advent-of-code-2022/solvers"
)

func main() {
	work_dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	solvers_map := map[string]func(string){
		"1": solvers.Solve1,
		"2": solvers.Solve2,
		"3": solvers.Solve3,
		"4": solvers.Solve4,
		"5": solvers.Solve5,
		"6": solvers.Solve6,
	}

	for i := 1; i <= 6; i++ {
		str_i := strconv.Itoa(i)
		file_path := work_dir + "/" + "inputs/" + str_i + ".txt"
		content := get_file_content(file_path)
		fmt.Println("\nSolving puzzle " + str_i + "...")
		solvers_map[str_i](content)
		fmt.Println("Solved puzzle " + str_i + "!")
	}
}

func get_file_content(file_path string) string {
	file_content, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}
	content := string(file_content)

	return content
}
