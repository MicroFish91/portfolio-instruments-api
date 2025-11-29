build:
	@go build -o bin/portfolio-instruments cmd/api/main.go 

run: build
	@./bin/portfolio-instruments	

test:
	@go test -count=1 -v ./...

migration:
	@migrate create -ext sql -dir migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down 

tc-redact:
	@go run cmd/tc_redact/main.go

config-up:
	@docker-compose up -d

config-down:
	@docker-compose down

pg-dump:
	@go run cmd/pg_commands/main.go dump

pg-restore:
	@go run cmd/pg_commands/main.go restore