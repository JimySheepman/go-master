default: build

.PHONY: lint
# analysis linting based on defined linting rules
lint:
	command -v golangci-lint || curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GOPATH}/bin latest
	golangci-lint run ./... -v

.PHONY: run
# runs mongo application
run:
	go run ./cmd/mongo/main.go

# runs mysql application
run-mysql:
	go run ./cmd/mysql/main.go

# runs postgreSQL application
run-postgres:
	go run ./cmd/postgres/main.go

.PHONY: vendor
# creates or updates project vendor, by downloading missing dependencies
vendor:
	go mod vendor

.PHONY: tidy
# It adds any missing module requirements necessary to build the current module’s packages and dependencies, and it
# removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to go.sum
# and removes unnecessary entries.
tidy:
	go mod tidy

.PHONY: build
# compiles existing programs, and creates its binaries into bin folder
build:
	go build -o bin/mongo ./cmd/mongo/main.go
	go build -o bin/mysql ./cmd/mysql/main.go
	go build -o bin/postgres ./cmd/postgres/main.go

.PHONY: generate
# runs go generate command, specially used to generate mocks
generate:
	go generate ./...

.PHONY: test
# runs tests
test:
	go test ./...

# runs coverage tests and generates the coverage report
test-coverage:
	go test ./... -v -coverpkg=./...

# runs integration tests
test-integration:
	go test ./... -tags=integration ./...

# clears generated mocks
clear-mock:
	find . -iname '*_mock.go' -exec rm {} \;

.PHONY: swagger
# generate swagger yml specs
swagger:
	swagger generate spec -o ./swagger.yaml --scan-models