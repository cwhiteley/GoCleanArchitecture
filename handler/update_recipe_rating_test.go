package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"recipes/appcontext"
	"recipes/domain"
	"recipes/repository"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateRecipeRatingWhenSuccessful(t *testing.T) {
	appcontext.Initialize()
	id := int64(1987236498)
	updatedRating := 3
	recipe := &domain.Recipe{Name: "salad",
		PreparationTime:  15,
		DifficultyRating: 1,
		IsVegetarian:     true,
		ID:               id}
	updatedRecipe := &domain.Recipe{Name: "salad",
		PreparationTime:  15,
		DifficultyRating: updatedRating,
		IsVegetarian:     true,
		ID:               id}

	mockedRecipeRepository := &repository.MockedRecipeRepository{}
	mockedRecipeRepository.On("FindByID", id).Return(recipe, nil)
	mockedRecipeRepository.On("Update", updatedRecipe).Return(nil)

	router := mux.NewRouter()
	router.HandleFunc("/v1/recipes/{id}/rating", UpdateRecipeRatingHandler(mockedRecipeRepository)).Methods("POST")
	ts := httptest.NewServer(router)
	defer ts.Close()

	rating := struct{ DifficultyRating int }{DifficultyRating: updatedRating}
	encodedRating, err := json.Marshal(rating)
	require.NoError(t, err, "error in encoding to JSON")
	r, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/recipes/%d/rating", ts.URL, id), bytes.NewReader(encodedRating))
	require.NoError(t, err, "failed to create a request")
	r.Header.Set("Content-Type", "application/json")

	response, err := ts.Client().Do(r)
	require.NoError(t, err, "failed to fire request")

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	returnedRecipe, err := readRecipeFromBody(response.Body)
	require.NoError(t, err, "Failed to unmarshal response")
	assert.Equal(t, updatedRecipe, returnedRecipe)
	mockedRecipeRepository.AssertExpectations(t)
}
