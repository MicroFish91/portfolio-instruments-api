build:
	@go build -o bin/portfolioinstruments cmd/api/main.go 

run: build
	@./bin/portfolioinstruments	

test:
	@go test -count=1 -v ./...

migration:
	@migrate create -ext sql -dir migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up 

migrate-down:
	@go run cmd/migrate/main.go down 

tc-scrub:
	@go run cmd/tc_scrub/main.go

config-up:
	@docker-compose up

config-down:
	@docker-compose down