package server

import (
	"recipes/repository"
)

type Dependencies struct {
	repository.RecipeStorer
}
