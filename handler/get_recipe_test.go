package handler

import (
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

func TestGetRecipeWhenSuccessful(t *testing.T) {
	appcontext.Initialize()
	id := int64(1987236498)
	recipe := &domain.Recipe{Name: "salad",
		PreparationTime:  15,
		DifficultyRating: 1,
		IsVegetarian:     true,
		ID:               id}

	mockedRecipeRepository := &repository.MockedRecipeRepository{}
	mockedRecipeRepository.On("FindByID", id).Return(recipe, nil)

	router := mux.NewRouter()
	router.HandleFunc("/v1/recipes/{id}", GetRecipeHandler(mockedRecipeRepository)).Methods("GET")
	ts := httptest.NewServer(router)
	defer ts.Close()

	r, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/recipes/%d", ts.URL, id), nil)
	require.NoError(t, err, "failed to create a request")
	r.Header.Set("Content-Type", "application/json")

	response, err := ts.Client().Do(r)
	require.NoError(t, err, "failed to fire request")

	assert.Equal(t, http.StatusOK, response.StatusCode)
	returnedRecipe, err := readRecipeFromBody(response.Body)
	require.NoError(t, err, "Failed to unmarshal response")
	assert.Equal(t, recipe, returnedRecipe)
	mockedRecipeRepository.AssertExpectations(t)
}
