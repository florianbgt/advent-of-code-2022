package solvers

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve4(content string) {
	pairs := strings.Split(content, "\n")

	list_of_pairs := get_list_of_pairs(pairs)

	count_duplicated := get_duplicate_count(list_of_pairs)
	count_overlaped := get_overlaped_count(list_of_pairs)

	fmt.Println("There is " + strconv.Itoa(count_duplicated) + " duplicates")
	fmt.Println("There is " + strconv.Itoa(count_overlaped) + " overlaps")
}

func get_list_of_pairs(pairs []string) [][]string {
	list_of_pairs := make([][]string, 0)

	for _, pair := range pairs {
		pair := strings.Split(pair, ",")
		list_of_pairs = append(list_of_pairs, pair)
	}

	return list_of_pairs
}

func create_list_of_items(pair string) []int {
	items := strings.Split(pair, string('-'))

	first, err := strconv.Atoi(items[0])
	if err != nil {
		panic(err)
	}

	last, err := strconv.Atoi(items[1])
	if err != nil {
		panic(err)
	}

	list := make([]int, 0)

	for i := first; i <= last; i++ {
		list = append(list, i)
	}

	return list
}

func remove_duplicates(list []int, target_list []int) []int {
	list_copy := make([]int, 0)
	for _, num := range list {
		list_copy = append(list_copy, num)
	}

	for _, target_list_num := range target_list {
		for list_i, list_num := range list_copy {
			if target_list_num == list_num {
				list_copy = append(list_copy[:list_i], list_copy[list_i+1:]...)
			}
		}
	}

	return list_copy
}

func is_contained(list1 []int, list2 []int) bool {
	list1_copy := remove_duplicates(list1, list2)
	list2_copy := remove_duplicates(list2, list1)

	if len(list1_copy) == 0 || len(list2_copy) == 0 {
		return true
	}

	return false
}

func get_duplicate_count(list_of_pairs [][]string) int {
	count := 0

	for _, pair := range list_of_pairs {
		first_pair := pair[0]
		second_pair := pair[1]

		first_list := create_list_of_items(first_pair)
		second_list := create_list_of_items(second_pair)

		is_contained := is_contained(first_list, second_list)

		if is_contained {
			count += 1
		}

	}

	return count
}

func is_overlap(list1 []int, list2 []int) bool {
	for _, list1_num := range list1 {
		for _, list2_num := range list2 {
			if list1_num == list2_num {
				return true
			}
		}
	}

	return false
}

func get_overlaped_count(list_of_pairs [][]string) int {
	count := 0

	for _, pair := range list_of_pairs {
		first_pair := pair[0]
		second_pair := pair[1]

		first_list := create_list_of_items(first_pair)
		second_list := create_list_of_items(second_pair)

		is_overlap := is_overlap(first_list, second_list)

		if is_overlap {
			count += 1
		}
	}
	return count
}
