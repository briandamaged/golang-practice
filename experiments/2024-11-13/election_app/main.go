package main

import (
	"fmt"

	"github.com/briandamaged/golang-practice/experiments/2024-11-13/election_lib"
)

func main() {
	var ranking = election_lib.Ranking{
		"foo",
		"bar",
	}

	preference, err := election_lib.GetPreference(ranking)
	if err == nil {
		fmt.Printf("Preference: %s\n", preference)
	} else {
		fmt.Println("Error: ", err)
	}
}
