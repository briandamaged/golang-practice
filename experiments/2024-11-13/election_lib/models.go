package election_lib

import (
	"errors"
)

type ID string
type Ranking []ID

func GetPreference(ranking Ranking) (ID, error) {
	if len(ranking) == 0 {
		return "", errors.New("all choices were eliminated")
	} else {
		return ranking[0], nil
	}
}
