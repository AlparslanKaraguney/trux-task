#!/bin/bash

# Generate package list excluding 'mocks'
PACKAGES=$(go list ./... | grep -v /mocks | grep -v /proto)

# Convert package list to comma-separated string
COVERPKG=$(echo $PACKAGES | tr ' ' ',')

# Run tests with coverage
go test -coverpkg="$COVERPKG" -coverprofile=coverage.out $PACKAGES

# Optional: Display total coverage percentage
COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}')
echo "Total Coverage: $COVERAGE"

go tool cover -html=coverage.out -o coverage.html

# Set required coverage threshold
REQUIRED_COVERAGE=70.0

# Use awk for floating-point comparison
awk -v cov="$COVERAGE" -v req="$REQUIRED_COVERAGE" 'BEGIN {
    if (cov+0 < req+0) {
        printf "Coverage is below %.1f%%. Current coverage: %.1f%%\n", req, cov;
        exit 1;
    } else {
        printf "Coverage meets the required threshold: %.1f%%\n", cov;
    }
}'
