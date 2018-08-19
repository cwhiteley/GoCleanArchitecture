package domain

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsValidReturnsFalseIfDifficultyIsGreaterThan3(t *testing.T) {
	recipe := Recipe{DifficultyRating: 4, PreparationTime: 30}
	assert.False(t, recipe.IsValid())
}

func TestIsValidReturnsFalseIfDifficultyIsLessThan1(t *testing.T) {
	recipe := Recipe{DifficultyRating: 0, PreparationTime: 30}
	assert.False(t, recipe.IsValid())
}

func TestIsValidReturnsTrueIfDifficultyIsLessThan4AndGreaterThan0(t *testing.T) {
	recipe := Recipe{DifficultyRating: 2, PreparationTime: 30}
	assert.True(t, recipe.IsValid())
}

func TestIsValidReturnsFalseIfPrepTimeIsNotPositive(t *testing.T) {
	recipe := Recipe{DifficultyRating: 2}
	assert.False(t, recipe.IsValid())
}

func TestIsValidReturnsFalseIfPrepTimeIsGreaterThanMaxLimit(t *testing.T) {
	maxPrepTime, err := strconv.Atoi(os.Getenv("MAX_PREP_TIME_IN_MINS"))
	require.NoError(t, err, "error converting max prep time to int")
	recipe := Recipe{DifficultyRating: 2, PreparationTime: maxPrepTime + 1}
	assert.False(t, recipe.IsValid())
}
