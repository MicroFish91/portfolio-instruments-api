build:
	@go build -o bin/portfolioinstruments cmd/main.go 

run: build
	@./bin/portfolioinstruments	

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up 

migrate-down:
	@go run cmd/migrate/main.go down 

config-up:
	@docker-compose up

config-down:
	@docker-compose down -v