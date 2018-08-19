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

func TestUpdateRecipeWhenSuccessful(t *testing.T) {
	appcontext.Initialize()
	id := int64(8912637948)
	updatedRecipe := &domain.Recipe{Name: "salad",
		PreparationTime:  15,
		DifficultyRating: 1,
		IsVegetarian:     true,
		ID:               id}
	encodedRecipe, err := json.Marshal(updatedRecipe)
	require.NoError(t, err, "failed to marshall recipe")

	mockedRecipeRepository := &repository.MockedRecipeRepository{}
	mockedRecipeRepository.On("Update", updatedRecipe).Return(nil)

	router := mux.NewRouter()
	router.HandleFunc("/v1/recipes/{id}", UpdateRecipeHandler(mockedRecipeRepository)).Methods("PUT")
	ts := httptest.NewServer(router)
	defer ts.Close()

	r, err := http.NewRequest("PUT", fmt.Sprintf("%s/v1/recipes/%d", ts.URL, id), bytes.NewReader(encodedRecipe))
	require.NoError(t, err, "failed to create a request")
	r.Header.Set("Content-Type", "application/json")

	response, err := ts.Client().Do(r)
	require.NoError(t, err, "failed to fire request")

	assert.Equal(t, http.StatusOK, response.StatusCode)
	returnedRecipe, err := readRecipeFromBody(response.Body)
	require.NoError(t, err, "Failed to unmarshal response")
	assert.Equal(t, updatedRecipe, returnedRecipe)
	mockedRecipeRepository.AssertExpectations(t)
}
