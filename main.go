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
		"1":  solvers.Solve1,
		"2":  solvers.Solve2,
		"3":  solvers.Solve3,
		"4":  solvers.Solve4,
		"5":  solvers.Solve5,
		"6":  solvers.Solve6,
		"7":  solvers.Solve7,
		"8":  solvers.Solve8,
		"9":  solvers.Solve9,
		"10": solvers.Solve10,
		"11": solvers.Solve11,
		"12": solvers.Solve12,
		"13": solvers.Solve13,
		"14": solvers.Solve14,
		"15": solvers.Solve15,
		"16": solvers.Solve16,
		"17": solvers.Solve17,
		"18": solvers.Solve18,
		"19": solvers.Solve19,
		"20": solvers.Solve20,
		"21": solvers.Solve21,
		"22": solvers.Solve22,
		"23": solvers.Solve23,
		"24": solvers.Solve24,
		"25": solvers.Solve25,
	}

	for i := 1; i <= 7; i++ {
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
