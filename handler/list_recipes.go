package handler

import (
	"database/sql"
	"net/http"

	"recipes/appcontext"
	"recipes/domain"
	"recipes/repository"
)

type RecipesList struct {
	Recipes []domain.Recipe `json:"recipes"`
}

func ListRecipesHandler(recipeRepository repository.RecipeStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		recipes, err := recipeRepository.FetchAll()
		if sql.ErrNoRows == err {
			appcontext.LogDebug("Recipes not found", "")
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			appcontext.LogError("Failed to read from DB", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		writeToBody(w, RecipesList{Recipes: recipes}, http.StatusOK)
		appcontext.LogInfo("Successfully fetched recipes", "")
	}
}
