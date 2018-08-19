package repository

import (
	"database/sql"
	"fmt"
	"testing"

	"recipes/appcontext"
	"recipes/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func withCleanDB(block func(), tableName string) {
	db := appcontext.GetDB()
	db.Exec(fmt.Sprintf("truncate %s", tableName))
	block()
	db.Exec(fmt.Sprintf("truncate %s", tableName))
}

func TestRecipeRepositoryWritesData(t *testing.T) {
	appcontext.Initialize()
	withCleanDB(func() {
		recipe := &domain.Recipe{Name: "test",
			PreparationTime:  30,
			DifficultyRating: 2,
			IsVegetarian:     true}
		repository := NewRecipeRepository()
		err := repository.Write(recipe)
		assert.NoError(t, err)

		persistedRecipe := &domain.Recipe{}
		db := appcontext.GetDB()
		err = db.Get(persistedRecipe, "select id, name, prep_time_in_minutes, difficulty, vegetarian from recipes;")
		require.NoError(t, err, "Error in fetching persisted recipe")
		assert.NotEqual(t, "", persistedRecipe.ID)
		assert.Equal(t, recipe.Name, persistedRecipe.Name)
		assert.Equal(t, recipe.PreparationTime, persistedRecipe.PreparationTime)
		assert.Equal(t, recipe.DifficultyRating, persistedRecipe.DifficultyRating)
		assert.Equal(t, recipe.IsVegetarian, persistedRecipe.IsVegetarian)
	}, "recipes")
}

func TestRecipeRepositoryReadsData(t *testing.T) {
	appcontext.Initialize()
	withCleanDB(func() {
		name := "foo"
		prepTime := 20
		difficulty := 3
		isVegetarian := false
		db := appcontext.GetDB()
		row, err := db.NamedQuery(writeRecipe,
			&domain.Recipe{Name: name,
				PreparationTime:  prepTime,
				DifficultyRating: difficulty,
				IsVegetarian:     isVegetarian})
		require.NoError(t, err, "Error in populating test data")
		var id int64
		if row.Next() {
			row.Scan(&id)
		}

		repository := NewRecipeRepository()
		recipe, err := repository.FindByID(id)
		assert.NoError(t, err)

		assert.Equal(t, id, recipe.ID)
		assert.Equal(t, name, recipe.Name)
		assert.Equal(t, prepTime, recipe.PreparationTime)
		assert.Equal(t, difficulty, recipe.DifficultyRating)
		assert.Equal(t, isVegetarian, recipe.IsVegetarian)
	}, "recipes")
}

func TestRecipeRepositoryUpdatesData(t *testing.T) {
	appcontext.Initialize()
	withCleanDB(func() {
		db := appcontext.GetDB()
		row, err := db.NamedQuery(writeRecipe,
			&domain.Recipe{Name: "foo",
				PreparationTime:  20,
				DifficultyRating: 3,
				IsVegetarian:     false})
		require.NoError(t, err, "Error in populating test data")
		var id int64
		if row.Next() {
			row.Scan(&id)
		}

		updatedRecipe := &domain.Recipe{Name: "bar",
			PreparationTime:  70,
			DifficultyRating: 2,
			IsVegetarian:     true,
			ID:               id}
		repository := NewRecipeRepository()
		err = repository.Update(updatedRecipe)
		assert.NoError(t, err)

		persistedRecipe := &domain.Recipe{}
		err = db.Get(persistedRecipe, "select id, name, prep_time_in_minutes, difficulty, vegetarian from recipes;")
		require.NoError(t, err, "Error in fetching persisted recipe")

		assert.Equal(t, updatedRecipe.ID, persistedRecipe.ID)
		assert.Equal(t, updatedRecipe.Name, persistedRecipe.Name)
		assert.Equal(t, updatedRecipe.PreparationTime, persistedRecipe.PreparationTime)
		assert.Equal(t, updatedRecipe.DifficultyRating, persistedRecipe.DifficultyRating)
		assert.Equal(t, updatedRecipe.IsVegetarian, persistedRecipe.IsVegetarian)
	}, "recipes")
}

func TestRecipeRepositoryDeletesData(t *testing.T) {
	appcontext.Initialize()
	withCleanDB(func() {
		db := appcontext.GetDB()
		row, err := db.NamedQuery(writeRecipe,
			&domain.Recipe{Name: "foo",
				PreparationTime:  20,
				DifficultyRating: 3,
				IsVegetarian:     false})
		require.NoError(t, err, "Error in populating test data")
		var id int64
		if row.Next() {
			row.Scan(&id)
		}

		repository := NewRecipeRepository()
		err = repository.Delete(id)
		assert.NoError(t, err)

		persistedRecipe := &domain.Recipe{}
		err = db.Get(persistedRecipe, "select id, name, prep_time_in_minutes, difficulty, vegetarian from recipes;")
		assert.Equal(t, sql.ErrNoRows, err)
	}, "recipes")
}

func TestRecipeRepositoryReadsAllRecipeData(t *testing.T) {
	appcontext.Initialize()
	withCleanDB(func() {
		recipe1 := &domain.Recipe{Name: "foo",
			PreparationTime:  20,
			DifficultyRating: 3,
			IsVegetarian:     false}
		db := appcontext.GetDB()
		row, err := db.NamedQuery(writeRecipe, recipe1)
		require.NoError(t, err, "Error in populating test data")
		var id int64
		if row.Next() {
			row.Scan(&id)
			recipe1.ID = id
		}

		recipe2 := &domain.Recipe{Name: "bar",
			PreparationTime:  89,
			DifficultyRating: 2,
			IsVegetarian:     false}
		row, err = db.NamedQuery(writeRecipe, recipe2)
		require.NoError(t, err, "Error in populating test data")
		if row.Next() {
			row.Scan(&id)
			recipe2.ID = id
		}

		repository := NewRecipeRepository()
		recipes, err := repository.FetchAll()
		assert.NoError(t, err)
		assert.Equal(t, 2, len(recipes))
		assert.Equal(t, *recipe1, recipes[0])
		assert.Equal(t, *recipe2, recipes[1])
	}, "recipes")
}

func TestRecipeRepositoryFindsRecipeGivenName(t *testing.T) {
	appcontext.Initialize()
	withCleanDB(func() {
		name := "foo"
		db := appcontext.GetDB()
		expectedRecipe := &domain.Recipe{Name: name,
			PreparationTime:  45,
			DifficultyRating: 2,
			IsVegetarian:     false}
		row, err := db.NamedQuery(writeRecipe, expectedRecipe)
		require.NoError(t, err, "Error in populating test data")
		var id int64
		if row.Next() {
			row.Scan(&id)
			expectedRecipe.ID = id
		}

		repository := NewRecipeRepository()
		recipe, err := repository.FindByName(name)
		assert.NoError(t, err)

		assert.Equal(t, expectedRecipe, recipe)
	}, "recipes")
}
