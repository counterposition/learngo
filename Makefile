build:
	@go build -o learn ./cmd/learn

	@echo "✅  Build successful"

clean:
	@go clean
	@rm -f learn

	@echo "✅  Cleaning successful"

test:
	@go test -v ./...

	@echo "✅  Test successful"

run: build
	@echo "Running..."

	@./learn

db:
	@createdb --owner learn learn 'Database for the project counterposition/learngo'

	@echo "✅  Created database 'learn'"

dropdb:
	@dropdb --if-exists learn

	@echo "✅  Dropped database 'learn'"

resetdb: dropdb db
