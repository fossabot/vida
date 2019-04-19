# migrations
.PHONY: db
db:
	- docker-compose up -d postgres

.PHONY: migrate-create
migrate-create:
	go run main.go migrate create $(name)

.PHONY: migrate-up
migrate-up: db
	go run main.go migrate up

.PHONY: migrate-down
migrate-down: db
	go run main.go migrate down

.PHONY: migrate-rollback
migrate-rollback: db
	go run main.go migrate rollback

.PHONY: test
test:
	go test -v ./...

.PHONY: test-ci
test-ci:
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: data
data:
	go run main.go data
