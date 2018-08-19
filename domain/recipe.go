package domain

import (
	"os"
	"strconv"
)

type Recipe struct {
	ID               int64  `db:"id"`
	Name             string `db:"name"`
	PreparationTime  int    `db:"prep_time_in_minutes"`
	DifficultyRating int    `db:"difficulty"`
	IsVegetarian     bool   `db:"vegetarian"`
}

func (recipe *Recipe) IsValid() bool {
	if recipe.isDifficultyRatingValid() && recipe.isPrepTimeValid() {
		return true
	}
	return false
}

func (recipe *Recipe) isPrepTimeValid() bool {
	maxPrepTime, _ := strconv.Atoi(os.Getenv("MAX_PREP_TIME_IN_MINS"))
	if recipe.PreparationTime < 1 || recipe.PreparationTime > maxPrepTime {
		return false
	}
	return true
}

func (recipe *Recipe) isDifficultyRatingValid() bool {
	if recipe.DifficultyRating > 3 || recipe.DifficultyRating < 1 {
		return false
	}
	return true
}
