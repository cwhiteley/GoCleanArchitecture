package handler

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"recipes/domain"
)

func extractIDFromMap(context map[string]string) (int64, error) {
	rawValue, ok := context["id"]
	if !ok {
		return 0, errors.New("ID is not present")
	}

	id, err := strconv.ParseInt(rawValue, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func readRecipeFromBody(requestBody io.ReadCloser) (*domain.Recipe, error) {
	var recipe domain.Recipe
	body, err := ioutil.ReadAll(requestBody)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &recipe)
	if err != nil {
		return nil, err
	}
	defer requestBody.Close()
	return &recipe, nil
}

func writeToBody(w http.ResponseWriter, content interface{}, statusCode int) error {
	jsonEncodedRecipe, err := json.MarshalIndent(content, "", " ")
	if err != nil {
		return err
	}
	w.WriteHeader(statusCode)
	w.Write(jsonEncodedRecipe)
	return nil
}
