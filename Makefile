# Makefile to build the project
GO=go
LINT=golangci-lint
GOSEC=gosec

COVERAGE = -coverprofile=coverage.txt -covermode=atomic

all: tidy test lint
travis-ci: tidy test-cov lint scan-gosec

test:
	${GO} test `${GO} list ./...`

test-cov:
	${GO} test `${GO} list ./...` ${COVERAGE}

test-int:
	${GO} test `${GO} list ./...` -tags=integration

test-int-cov:
	${GO} test `${GO} list ./...` -tags=integration ${COVERAGE}

# Main lint: force our config, run single-threaded, keep it fast to avoid OOM on CI
lint:
	GOGC=50 ${LINT} run --config .golangci.yml --build-tags=integration,examples -j 1 --timeout 7m --fast

# Fallback: run linters separately if the combined run gets killed on CI
lint-split:
	GOGC=50 ${LINT} run --config .golangci.yml -j 1 --timeout 7m --fast --disable-all -E errcheck
	GOGC=50 ${LINT} run --config .golangci.yml -j 1 --timeout 7m --fast --disable-all -E staticcheck

scan-gosec:
	${GOSEC} ./...

tidy:
	${GO} mod tidy
