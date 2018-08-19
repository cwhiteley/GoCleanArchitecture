package handler

import (
	"net/http"

	"recipes/appcontext"
	"recipes/repository"
)

func CreateRecipeHandler(recipeRepository repository.RecipeStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		recipe, err := readRecipeFromBody(r.Body)
		if err != nil {
			appcontext.LogError("Failed to read body", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !recipe.IsValid() {
			appcontext.LogDebug("Invalid input", "")
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		err = recipeRepository.Write(recipe)
		if err != nil {
			appcontext.LogError("Failed to write to DB", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = writeToBody(w, recipe, http.StatusCreated)
		if err != nil {
			appcontext.LogError("Failed to serialize to JSON", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		appcontext.LogInfo("Successfully created recipe", string(recipe.ID))
	}
}
