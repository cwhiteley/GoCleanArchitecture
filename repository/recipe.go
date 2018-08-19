package repository

import (
	"recipes/appcontext"
	"recipes/domain"

	"github.com/jmoiron/sqlx"
)

const (
	writeRecipe = `insert into recipes (name, prep_time_in_minutes, difficulty, vegetarian)
										values (:name, :prep_time_in_minutes, :difficulty, :vegetarian) returning id`
	findRecipeByID = "select * from recipes where id = $1"
	updateRecipe   = `update recipes set name = :name, prep_time_in_minutes = :prep_time_in_minutes,
										difficulty = :difficulty, vegetarian = :vegetarian where id = :id`
	deleteRecipe     = "delete from recipes where id = $1"
	fetchAll         = "select * from recipes"
	findRecipeByName = "select * from recipes where name = $1"
)

type RecipeStorer interface {
	Write(*domain.Recipe) error
	FindByID(int64) (*domain.Recipe, error)
	Update(*domain.Recipe) error
	Delete(int64) error
	FetchAll() ([]domain.Recipe, error)
	FindByName(string) (*domain.Recipe, error)
}

type RecipeRepository struct {
	db *sqlx.DB
}

func NewRecipeRepository() RecipeRepository {
	return RecipeRepository{
		db: appcontext.GetDB(),
	}
}

func (rep RecipeRepository) Write(recipe *domain.Recipe) error {
	rows, err := rep.db.NamedQuery(writeRecipe, recipe)
	if err != nil {
		return err
	}

	var id int64
	if rows.Next() {
		rows.Scan(&id)
		recipe.ID = id
	}

	return nil
}

func (rep RecipeRepository) FindByID(id int64) (*domain.Recipe, error) {
	recipe := domain.Recipe{}
	err := rep.db.Get(&recipe, findRecipeByID, id)
	if err != nil {
		return nil, err
	}

	return &recipe, nil
}

func (rep RecipeRepository) Update(recipe *domain.Recipe) error {
	_, err := rep.db.NamedQuery(updateRecipe, recipe)
	return err
}

func (rep RecipeRepository) Delete(id int64) error {
	_, err := rep.db.Exec(deleteRecipe, id)
	return err
}

func (rep RecipeRepository) FetchAll() ([]domain.Recipe, error) {
	var recipes []domain.Recipe
	err := rep.db.Select(&recipes, fetchAll)
	return recipes, err
}

func (rep RecipeRepository) FindByName(name string) (*domain.Recipe, error) {
	recipe := domain.Recipe{}
	err := rep.db.Get(&recipe, findRecipeByName, name)
	if err != nil {
		return nil, err
	}

	return &recipe, nil
}
