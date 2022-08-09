#!/bin/bash

MODULE_NAME=""
TESTCASE_NAME=""
COVERAGE_REPORT="out/cover.out"
COVERAGE_MODE="count"
### Mod
go mod init ${MODULE_NAME}
go mod tidy
go get -u ${MODULE_NAME}
### Testing

go test -run ${TESTCASE_NAME}
go test -v ./...    # Test all under current package
go test -covermode=${COVERAGE_MODE}  -coverprofile=${COVERAGE_REPORT} ./... # Generate test coverage report

### Build


### Tool
go tool cover -func out/cover.out # View test coverage report by func
go tool compile -help

### Trace
go tool trace <trace_file> // trace file generated by -trace=trace.out in go test cmd