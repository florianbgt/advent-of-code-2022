package solvers

import (
	"fmt"
	"strconv"
)

func Solve6(content string) {
	message := content

	i := get_first_duplicate_index(4, message)

	j := get_first_duplicate_index(14, message)

	fmt.Println("first duplicate happen at character " + strconv.Itoa(i))
	fmt.Println("first duplicate happen at character " + strconv.Itoa(j))
}

func get_first_duplicate_index(num int, message string) int {
	for i := num - 1; i < len(message); i++ {
		list_chars := make([]byte, 0)
		for j := num - 1; j >= 0; j-- {
			list_chars = append(list_chars, message[i-j])
		}

		chars_map := make(map[byte]bool, 0)
		has_duplicates := false

		for _, char := range list_chars {
			if chars_map[char] == true {
				has_duplicates = true
			}
			chars_map[char] = true
		}

		if has_duplicates == false {
			return i + 1
		}

	}
	return 0
}
