build:
	@go build -o bin/SmartParkingSystem main.go

test:
	@go test -v ./...

run:build
	@./bin/SmartParkingSystem

migration:
	@migrate create -ext sql -dir db/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/migrate.go up

migrate-down:
	@go run cmd/migrate/migrate.go down

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down

docker-image-clean:
	docker image prune -a
	docker images -f "dangling=true" -q | xargs -r docker rmi
