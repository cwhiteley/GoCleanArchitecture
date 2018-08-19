package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"recipes/appcontext"
	"recipes/domain"
	"recipes/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateRecipeWhenSuccessful(t *testing.T) {
	appcontext.Initialize()
	newRecipe := &domain.Recipe{Name: "salad",
		PreparationTime:  15,
		DifficultyRating: 1,
		IsVegetarian:     true}
	encodedRecipe, err := json.Marshal(newRecipe)
	require.NoError(t, err, "failed to marshall recipe")

	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/v1/recipe", bytes.NewReader(encodedRecipe))
	require.NoError(t, err, "failed to create a request")
	r.Header.Set("Content-Type", "application/json")

	mockedRecipeRepository := &repository.MockedRecipeRepository{}
	mockedRecipeRepository.On("Write", newRecipe).Return(nil)

	CreateRecipeHandler(mockedRecipeRepository)(w, r)

	returnedRecipe := &domain.Recipe{}
	err = json.Unmarshal(w.Body.Bytes(), returnedRecipe)
	require.NoError(t, err, "Failed to unmarshal response")
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, newRecipe, returnedRecipe)
	mockedRecipeRepository.AssertExpectations(t)
}
