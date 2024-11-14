package main

import (
	"fmt"

	"github.com/briandamaged/golang-practice/experiments/2024-11-13/election_lib/model"
)

func PrintOutcome(ranking model.TotalRanking[model.ID]) {
	preference, err := ranking.GetFavorite()
	if err == nil {
		fmt.Printf("Preference: %s\n", preference)
	} else {
		fmt.Println("Error: ", err)
	}
}

func main() {
	var ranking = model.TotalRanking[model.ID]{
		"foo",
		"bar",
		"quuz",
		"zerp",
	}

	PrintOutcome(ranking)
	ranking = ranking.Disqualify("foo")
	PrintOutcome(ranking)
	ranking = ranking.Disqualify("quuz")
	PrintOutcome(ranking)
}
