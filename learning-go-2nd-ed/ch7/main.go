package main

import (
	"io"
	"log"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (league League) GetTeam(n string) (*Team, bool) {
	for _, t := range league.Teams {
		if t.Name == n {
			return &t, true
		}
	}

	return nil, false
}

func (league League) GivePointTo(n string) {
	_, ok := league.GetTeam(n)
	if !ok {
		log.Fatalf("No such team: %s\n", n)
	}
	league.Wins[n]++
}

func (league League) MatchResult(n1 string, s1 int, n2 string, s2 int) {
	switch {
	case s1 > s2:
		league.GivePointTo(n1)
	case s1 < s2:
		league.GivePointTo(n2)
	}
}

func (league League) Ranking() []string {
	var retval = []string{}
	for _, t := range league.Teams {
		retval = append(retval, t.Name)
	}

	sort.Slice(retval, func(i, j int) bool {
		return league.Wins[retval[i]] > league.Wins[retval[j]]
	})

	return retval
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(ranker Ranker, w io.Writer) {
	var ranking = ranker.Ranking()
	for _, r := range ranking {
		io.WriteString(w, r)
		io.WriteString(w, "\n")
	}
}

func main() {
	var league = League{
		Teams: []Team{
			{
				Name: "one",
				Players: []string{
					"Babe Ruth",
					"Ken Griffey Jr.",
				},
			},
			{
				Name: "two",
				Players: []string{
					"Roger Maris",
				},
			},
			{
				Name:    "three",
				Players: []string{},
			},
		},
		Wins: map[string]int{},
	}

	league.MatchResult("one", 3, "two", 2)
	league.MatchResult("one", 3, "two", 2)
	league.MatchResult("one", 3, "two", 5)
	league.MatchResult("one", 3, "two", 5)
	league.MatchResult("one", 3, "two", 5)
	RankPrinter(league, os.Stdout)
}
