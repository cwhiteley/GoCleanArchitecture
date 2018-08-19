package repository

import (
	"recipes/domain"

	"github.com/stretchr/testify/mock"
)

type MockedRecipeRepository struct {
	mock.Mock
}

func (mock MockedRecipeRepository) Write(recipe *domain.Recipe) error {
	args := mock.Called(recipe)
	return args.Error(0)
}

func (mock MockedRecipeRepository) FindByID(id int64) (*domain.Recipe, error) {
	args := mock.Called(id)
	return args[0].(*domain.Recipe), args.Error(1)
}

func (mock MockedRecipeRepository) Update(recipe *domain.Recipe) error {
	args := mock.Called(recipe)
	return args.Error(0)
}

func (mock MockedRecipeRepository) Delete(id int64) error {
	args := mock.Called(id)
	return args.Error(0)
}

func (mock MockedRecipeRepository) FetchAll() ([]domain.Recipe, error) {
	args := mock.Called()
	return args[0].([]domain.Recipe), args.Error(1)
}

func (mock MockedRecipeRepository) FindByName(name string) (*domain.Recipe, error) {
	args := mock.Called(name)
	return args[0].(*domain.Recipe), args.Error(1)
}
