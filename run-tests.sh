#!/bin/bash

# run-tests.sh
echo "Running tests in Docker..."

# Build and run tests
docker-compose run --rm go-test

# Check exit code
if [ $? -eq 0 ]; then
    echo "Tests passed successfully!"
else
    echo "Tests failed!"
    exit 1
fi