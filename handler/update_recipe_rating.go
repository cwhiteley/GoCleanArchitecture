package handler

import (
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"recipes/appcontext"
	"recipes/repository"

	"github.com/gorilla/mux"
)

func UpdateRecipeRatingHandler(recipeRepository repository.RecipeStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, err := extractIDFromMap(mux.Vars(r))
		if err != nil {
			appcontext.LogError("Failed to read id", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		rating, err := readRating(r.Body)
		if err != nil {
			appcontext.LogError("Failed to read from body", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		recipe, err := recipeRepository.FindByID(id)
		if sql.ErrNoRows == err {
			appcontext.LogDebug("No recipe found", string(id))
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			appcontext.LogError("Failed to read from DB", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		recipe.DifficultyRating = rating
		if !recipe.IsValid() {
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

		writeToBody(w, recipe, http.StatusCreated)
		appcontext.LogInfo("Successfully updated recipe", string(recipe.ID))
	}
}

func readRating(requestBody io.ReadCloser) (int, error) {
	var rating struct{ DifficultyRating int }
	body, err := ioutil.ReadAll(requestBody)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(body, &rating)
	if err != nil {
		return 0, err
	}
	defer requestBody.Close()
	return rating.DifficultyRating, nil
}
