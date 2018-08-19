package handler

import (
	"net/http"

	"recipes/appcontext"
	"recipes/repository"

	"github.com/gorilla/mux"
)

func DeleteRecipeHandler(recipeRepository repository.RecipeStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, err := extractIDFromMap(mux.Vars(r))
		if err != nil {
			appcontext.LogError("Failed to read id", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = recipeRepository.Delete(id)
		if err != nil {
			appcontext.LogError("Failed to delete from DB", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		appcontext.LogInfo("Successfully deleted recipe", string(id))
	}
}
