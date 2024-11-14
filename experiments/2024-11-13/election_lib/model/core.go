package model

import (
	"errors"
)

type ID string

type Poll struct {
	Id      ID
	Desc    string
	Choices []Choice
}

type Choice struct {
	Id    ID
	Label string
}

type TotalRanking[T comparable] []T

func (r TotalRanking[T]) IsEmpty() bool {
	return len(r) == 0
}

func (r TotalRanking[T]) HasFavorite() bool {
	return len(r) > 0
}

func (r TotalRanking[T]) GetFavorite() (T, error) {
	if r.HasFavorite() {
		return r[0], nil
	} else {
		var zero T
		return zero, errors.New("TotalRanking is empty")
	}
}

func (r TotalRanking[T]) Disqualify(id T) TotalRanking[T] {
	var retval TotalRanking[T] = make(TotalRanking[T], 0, len(r))
	for _, item := range r {
		if item != id {
			retval = append(retval, item)
		}
	}
	return retval
}
