# Go-Webapp

## Running the project

`docker-compose up`

Then access the server through the exposed by docker port `8000`

## Package management

To remove unused downloaded packages

`go mod tidy`

## Testing

To run all test files in the current directory (with verbose flag)

`go test -v`

To see test coverage

`go test -cover`

To see more details on the test coverage as an HTML file

`go test -coverprofile=coverage.out && go tool cover -html=coverage.out`