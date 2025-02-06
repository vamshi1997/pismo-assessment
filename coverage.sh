echo "Running tests with coverage..."
go test -covermode=atomic -coverprofile=coverage.out ./... 2>&1 | grep -v "no test files"

echo -e "\nCoverage by package:"
go tool cover -func=coverage.out | grep -v "total:" | \
    sed 's/\(.*\)\.go.*coverage: \(.*\) of statements$/\1\t\2/'

echo -e "\nTotal coverage:"
go tool cover -func=coverage.out | grep "total:" | \
    sed 's/total:\s*statements\s*\(.*\)$/\1/'