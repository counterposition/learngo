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
