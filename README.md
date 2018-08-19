# Clean Architecture on Golang
[Uncle Bob's The Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)

Description
After reading the uncle Bob’s Clean Architecture Concept, I’m trying to implement it in Golang using RESTful Recipes API.

Rule of Clean Architecture by Uncle Bob

* Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
* Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
* Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
* Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

This project has 4 layers:
* Models Layer: Domain
* Repository Layer: Interfaces
* Usecase Layer: Handlers
* Delivery Layer: Infrastructure

## Assumptions
* It's OK to actually delete a recipe from the DB in the context of this problem. But in production, when a recipe resource is delete, the data would be archived.
* The rate API ( POST /recipes/{id}/rating) updates the difficulty rating of the recipe.

## Tradeoffs made
* I have not spent got enough time to do TDD for the branches that deal with failures since the scaffolding also took time and we're not supposed to use frameworks. So test coverage is going to be less than 70%.
* I have tried to write as readable code as possible so that documentation is not required.
* Have added authentication middleware that only checks for the presence of the Authorization header. Ideally this middleware should verify the access token by make a request to an upstream authorizing server or central redis for validation.
* The recipe search API does an exact name match and no fuzzy search. Ideally it should have fuzzy matching at DB level or application level. To handle even more traffic, we can index all the recipes in Elastic Search and define must/should filters for our fuzzy matcher.

## TODOs
* Using jwt token for authentication and storing the token in redis
* Creating domain for jwt token:
    ```go
    // Authentication struct
    type Authentication struct {
    	Username string `json:"username,omitempty"`
    	Password string `json:"password,omitempty"`
    }

    // Jwt struct
    type Jwt struct {
    	Token string `json:"token"`
    }

    // TokenClaims struct
    type TokenClaims struct {
    	Username string    `json:"username"`
    	Time     time.Time `json:"time"`
    	UserID   int       `json:"userid"`
    }
    ```
* Adding RatingRequest domain and a separate table/array field for storing the ratings
    ```go
    type RatingRequest struct {
    	RecipeID int
    	Ratings  int
        Time     time.Time
    }
    ```
* Writing custom middleware or using negroni
* Caching recipes in redis to make the read operations faster.
* Extract config as a package so that each config has a type. We should also check to see if config is present.
* Adding pagination for get all recipes API
* Using secure middleware for quick wins in security

## Get Started(Make)
* Modify `.env.sample` with relevant environment variables. Run `source .env.sample` to populate environment variables needed to run the app.
* Run `make build`.
* Run `make test` to run test.
* Run `make fmt` to run go-fmt.
* Running `make` builds, runs tests and lints the codebase.

## Get Started(Go Build)
* Create the go project workspace
    `mkdir go/src/kumar-sa-api-test`
    `export GOPATH=go/src/kumar-sa-api-test`
* Build the go project
    `go build kumar-sa-api-test`

## Get Started(Docker)
* Running `docker-compose up -d` to spin up the docker containers.

## APIs

#### Create Recipe
```json
POST /v1/recipe

{
    "name": "bar",
    "prep_time_in_minutes":  70,
    "difficulty": 2,
    "vegetarian":     true,
}
```

#### GET Recipe
```json
GET /v1/recipe/{id}
```

#### List Recipe
```json
GET /v1/recipes
```

#### Update Recipe
```json
PUT /v1/recipe/{id}

{
    "name": "bar",
    "prep_time_in_minutes":  90,
    "difficulty": 2,
    "vegetarian":     true,
}
```

#### Delete Recipe
```json
DELETE /v1/recipe/{id}
```

### Delete Recipe
```json
DELETE /v1/recipe/{id}
```

### Search Recipe by Name
```json
GET /v1/recipe?name=foo
```

### Update Recipe Rating
```json
POST /v1/recipes/{id}/rating
```
