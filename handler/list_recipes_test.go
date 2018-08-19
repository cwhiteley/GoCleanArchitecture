package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func TestListRecipesWhenSuccessful(t *testing.T) {
	appcontext.Initialize()
	recipe1 := domain.Recipe{Name: "salad",
		PreparationTime:  15,
		DifficultyRating: 1,
		IsVegetarian:     true,
		ID:               int64(1987236498)}
	recipe2 := domain.Recipe{Name: "bread",
		PreparationTime:  75,
		DifficultyRating: 3,
		IsVegetarian:     true,
		ID:               int64(726367263)}
	recipes := []domain.Recipe{recipe1, recipe2}

	mockedRecipeRepository := &repository.MockedRecipeRepository{}
	mockedRecipeRepository.On("FetchAll").Return(recipes, nil)

	router := mux.NewRouter()
	router.HandleFunc("/v1/recipes", ListRecipesHandler(mockedRecipeRepository)).Methods("GET")
	ts := httptest.NewServer(router)
	defer ts.Close()

	r, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/recipes", ts.URL), nil)
	require.NoError(t, err, "failed to create a request")
	r.Header.Set("Content-Type", "application/json")

	response, err := ts.Client().Do(r)
	require.NoError(t, err, "failed to fire request")

	assert.Equal(t, http.StatusOK, response.StatusCode)

	var returnedRecipes RecipesList
	body, err := ioutil.ReadAll(response.Body)
	require.NoError(t, err, "failed to read list recipes response")
	err = json.Unmarshal(body, &returnedRecipes)
	require.NoError(t, err, "failed to unmarshal list recipes response")
	defer response.Body.Close()
	require.NoError(t, err, "Failed to unmarshal response")
	assert.Equal(t, RecipesList{Recipes: recipes}, returnedRecipes)
	mockedRecipeRepository.AssertExpectations(t)
}
