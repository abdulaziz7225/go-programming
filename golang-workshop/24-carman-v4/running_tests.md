# Running testing and get coverage results

1. Run the tests and generate a coverage profile

    ```bash
    go test ./internal/http -v -coverprofile=coverage.out
    ```

2. View detailed coverage in the CLI

    ```bash
    go tool cover -func=coverage.out
    ```

3. View the interactive HTML report

    ```bash
    go tool cover -html=coverage.out
    ```
