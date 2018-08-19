.PHONY: all
all: setup build fmt vet lint test

APP=recipes-app
GLIDE_NOVENDOR=$(shell glide novendor)
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
UNIT_TEST_PACKAGES=$(shell glide novendor | grep -v "featuretests")
APP_EXECUTABLE="./out/$(APP)"

setup:
	go get -u github.com/golang/lint/golint
	go get github.com/DATA-DOG/godog/cmd/godog

build-deps:
	glide install

update-deps:
	glide update

compile:
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)

build: build-deps compile

install:
	go install ./...

fmt:
	go fmt $(GLIDE_NOVENDOR)

vet:
	go vet $(GLIDE_NOVENDOR)

lint:
	@for p in $(UNIT_TEST_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } \
	done

test: db.migrate
	ENVIRONMENT=test go test $(UNIT_TEST_PACKAGES) -p=1

db.migrate:
	$(APP_EXECUTABLE) migrate

db.rollback:
	$(APP_EXECUTABLE) rollback

run:
	$(APP_EXECUTABLE) start