# veriff-api-client-go

[![Go Reference](https://pkg.go.dev/badge/github.com/brokeyourbike/veriff-api-client-go.svg)](https://pkg.go.dev/github.com/brokeyourbike/veriff-api-client-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/brokeyourbike/veriff-api-client-go)](https://goreportcard.com/report/github.com/brokeyourbike/veriff-api-client-go)
[![Maintainability](https://api.codeclimate.com/v1/badges/698070207d9d5a7bb44e/maintainability)](https://codeclimate.com/github/brokeyourbike/veriff-api-client-go/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/698070207d9d5a7bb44e/test_coverage)](https://codeclimate.com/github/brokeyourbike/veriff-api-client-go/test_coverage)

Veriff API Client for Go

## Installation

```bash
go get github.com/brokeyourbike/veriff-api-client-go
```

## Usage

```go
client := veriff.NewClient("veriff.com", "token")
client.CreateSession(context.TODO(), veriff.CreateSessionPayload{})
```

## Authors
- [Ivan Stasiuk](https://github.com/brokeyourbike) | [Twitter](https://twitter.com/brokeyourbike) | [LinkedIn](https://www.linkedin.com/in/brokeyourbike) | [stasi.uk](https://stasi.uk)

## License
[BSD-3-Clause License](https://github.com/brokeyourbike/veriff-api-client-go/blob/main/LICENSE)