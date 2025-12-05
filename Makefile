include .env

proto-format:
	buf format -w

proto-lint:
	buf lint

proto-gen:
	buf generate

MOCK_GEN=go run go.uber.org/mock/mockgen@v0.6.0

mock-gen:
	$(MOCK_GEN) -source=internal/domain/element/element.go -destination=mocks/domain/element/mock_element.go -package=mocks
	$(MOCK_GEN) -source=internal/domain/element/repository.go -destination=mocks/domain/element/mock_repository.go -package=mocks
	$(MOCK_GEN) -source=internal/domain/category/category.go -destination=mocks/domain/category/mock_category.go -package=mocks
	$(MOCK_GEN) -source=internal/domain/category/repository.go -destination=mocks/domain/category/mock_repository.go -package=mocks
	$(MOCK_GEN) -source=internal/application/combination/usecase.go -destination=mocks/application/combination/mock_usecase.go -package=mocks

MIGRATE=migrate -source file://internal/infrastructure/db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_HOST_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)"

migrate-up:
	$(MIGRATE) up

migrate-up-one:
	$(MIGRATE) up 1

migrate-down:
	$(MIGRATE) down 1

migrate-reset:
	$(MIGRATE) drop -f

migrate-create:
	migrate create -ext sql -dir internal/infrastructure/db/migrations -format 20060102150405 $(name)

migrate-status:
	$(MIGRATE) version

test:
	mkdir -p tmp
	go test -cover ./internal/... -coverprofile=./tmp/cover.out
	go tool cover -func=./tmp/cover.out | tail -n 1
	go tool cover -html=./tmp/cover.out -o ./tmp/cover.html
	open ./tmp/cover.html