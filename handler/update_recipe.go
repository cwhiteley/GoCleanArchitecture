package handler

import (
	"net/http"

	"recipes/appcontext"
	"recipes/repository"

	"github.com/gorilla/mux"
)

func UpdateRecipeHandler(recipeRepository repository.RecipeStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, err := extractIDFromMap(mux.Vars(r))
		if err != nil {
			appcontext.LogError("Failed to read id", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		recipe, err := readRecipeFromBody(r.Body)
		if err != nil {
			appcontext.LogError("Failed to read body", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if recipe.ID != id || !recipe.IsValid() {
			appcontext.LogDebug("Invalid input", string(id))
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		err = recipeRepository.Update(recipe)
		if err != nil {
			appcontext.LogError("Failed to write to DB", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		writeToBody(w, recipe, http.StatusOK)
		appcontext.LogInfo("Successfully updated recipe", string(recipe.ID))
	}
}
