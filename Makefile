build:
	@go build -o learn

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

generate:
	@go run github.com/99designs/gqlgen generate idl/schema.graphql

	@echo "✅  Generate successful"
