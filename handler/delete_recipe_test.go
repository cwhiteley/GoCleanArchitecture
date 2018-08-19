package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"recipes/appcontext"
	"recipes/repository"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteRecipeWhenSuccessful(t *testing.T) {
	appcontext.Initialize()
	id := int64(1987236498)

	mockedRecipeRepository := &repository.MockedRecipeRepository{}
	mockedRecipeRepository.On("Delete", id).Return(nil)

	router := mux.NewRouter()
	router.HandleFunc("/v1/recipes/{id}", DeleteRecipeHandler(mockedRecipeRepository)).Methods("DELETE")
	ts := httptest.NewServer(router)
	defer ts.Close()

	r, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v1/recipes/%d", ts.URL, id), nil)
	require.NoError(t, err, "failed to create a request")
	r.Header.Set("Content-Type", "application/json")

	response, err := ts.Client().Do(r)
	require.NoError(t, err, "failed to fire request")

	assert.Equal(t, http.StatusOK, response.StatusCode)
	mockedRecipeRepository.AssertExpectations(t)
}
