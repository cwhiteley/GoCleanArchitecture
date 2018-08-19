package handler

import (
	"database/sql"
	"net/http"

	"recipes/appcontext"
	"recipes/repository"
)

func SearchRecipeHandler(recipeRepository repository.RecipeStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		name := r.URL.Query().Get("name")
		if name == "" {
			appcontext.LogDebug("Name is not present", "")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		recipe, err := recipeRepository.FindByName(name)
		if sql.ErrNoRows == err {
			appcontext.LogDebug("Recipe not found", name)
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			appcontext.LogError("Failed to read from DB", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		writeToBody(w, recipe, http.StatusOK)
		appcontext.LogInfo("Successfully fetched recipe", name)
	}
}
