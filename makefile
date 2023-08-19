.PHONY: default run build test docs clean run-with-docs
# Variables
APP_NAME=gopportunities

# Tasks
default: run-with-docs

run:
	@go run cmd/easyInvest/main.go
run-with-docs:
	@swag init
	@go run main.go
build: 
	@go build -o $(APP_NAME) cmd/easyInvest/main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs
