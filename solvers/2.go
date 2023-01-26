package solvers

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve2(content string) {
	rounds := strings.Split(content, "\n")

	total_score_if_sign, total_score_if_outcome := get_total_score(rounds)

	fmt.Println("The total score is " + strconv.Itoa(total_score_if_sign) + " if sign.")
	fmt.Println("The total score is " + strconv.Itoa(total_score_if_outcome) + " if outcome.")
}

func get_score_from_sign(opponent_sign string, sign string) int {
	score_sign_map := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	score_event_map := map[string]int{
		"win":  6,
		"loss": 0,
		"draw": 3,
	}

	event_sign_map := map[string]map[string]string{
		"X": { // Rock
			"A": "draw", // Rock
			"B": "loss", // Paper
			"C": "win",  // Scissors
		},
		"Y": { // Paper
			"A": "win",  // Rock
			"B": "draw", // Paper
			"C": "loss", // Scissors
		},
		"Z": { // Scissors
			"A": "loss", // Rock
			"B": "win",  // Paper
			"C": "draw", // Scissors
		},
	}

	score_sign := score_sign_map[sign]
	event := event_sign_map[sign][opponent_sign]
	score_event := score_event_map[event]

	score := score_sign + score_event

	return score
}

func get_sign_from_outcome(opponent_sign string, outcome string) string {
	outcome_sign_map := map[string]map[string]string{
		"X": { // loss
			"A": "Z", // Rock vs Scissors
			"B": "X", // Paper vs Rock
			"C": "Y", // Scissors vs Paper
		},
		"Y": { // draw
			"A": "X", // Rock vs Rock
			"B": "Y", // Paper vs Paper
			"C": "Z", // Acissors vs Scissors
		},
		"Z": { // win
			"A": "Y", // Rock vs Paper
			"B": "Z", // Paper vs Scissors
			"C": "X", // Scissors vs Rock
		},
	}

	sign := outcome_sign_map[outcome][opponent_sign]

	return sign
}

func get_total_score(rounds []string) (int, int) {
	total_score_if_sign := 0
	total_score_if_outcome := 0

	for _, round := range rounds {
		signs := strings.Split(round, " ")
		opponent_sign := signs[0]
		sign := signs[1]

		score_if_sign := get_score_from_sign(opponent_sign, sign)

		outcome := sign
		sign_for_outcome := get_sign_from_outcome(opponent_sign, outcome)
		score_if_outcome := get_score_from_sign(opponent_sign, sign_for_outcome)

		total_score_if_sign += score_if_sign
		total_score_if_outcome += score_if_outcome
	}

	return total_score_if_sign, total_score_if_outcome
}
