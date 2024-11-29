#!/bin/bash

# Generate package list excluding 'mocks'
PACKAGES=$(go list ./... | grep -v /mocks | grep -v /proto)

# Convert package list to comma-separated string
COVERPKG=$(echo $PACKAGES | tr ' ' ',')

go clean -testcache

# Run tests with coverage
go test -coverpkg="$COVERPKG" -coverprofile=coverage.out $PACKAGES

# Optional: Display total coverage percentage
COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}')
echo "Total Coverage: $COVERAGE"

go tool cover -html=coverage.out -o coverage.html

