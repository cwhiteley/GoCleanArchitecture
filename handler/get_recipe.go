package handler

import (
	"database/sql"
	"net/http"

	"recipes/appcontext"
	"recipes/repository"

	"github.com/gorilla/mux"
)

func GetRecipeHandler(recipeRepository repository.RecipeStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, err := extractIDFromMap(mux.Vars(r))
		if err != nil {
			appcontext.LogError("Failed to read id", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		recipe, err := recipeRepository.FindByID(id)
		if sql.ErrNoRows == err {
			appcontext.LogDebug("Recipe not found", string(id))
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			appcontext.LogError("Failed to read from DB", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		writeToBody(w, recipe, http.StatusOK)
		appcontext.LogInfo("Successfully fetched recipe", string(id))
	}
}
