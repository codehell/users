user = postgres
password = secret
database = users

create-migration:
	migrate create -ext sql -dir db/migrations -seq $(arg)

migrate-up:
	migrate -database 'postgres://postgres:$(password)@localhost:5432/$(database)?sslmode=disable' -path db/migrations up $(arg)

migrate-down:
	migrate -database 'postgres://postgres:$(password)@localhost:5432/$(database)?sslmode=disable' -path db/migrations down $(arg)

migrate-version:
	migrate -database 'postgres://postgres:$(password)@localhost:5432/$(database)?sslmode=disable' -path db/migrations version
  
go-test:
	docker run --rm --interactive --tty \
  -e FIRESTORE_EMULATOR_HOST=localhost:8442 \
  --volume $(PWD):/app \
  -w=/app \
  golang:1.14-buster go test -v