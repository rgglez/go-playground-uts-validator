# go-playground-uts-validator

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![GitHub all releases](https://img.shields.io/github/downloads/rgglez/go-playground-uts-validator/total)
![GitHub issues](https://img.shields.io/github/issues/rgglez/go-playground-uts-validator)
![GitHub commit activity](https://img.shields.io/github/commit-activity/y/rgglez/go-playground-uts-validator)
[![Go Report Card](https://goreportcard.com/badge/github.com/rgglez/go-playground-uts-validator)](https://goreportcard.com/report/github.com/rgglez/go-playground-uts-validator)
[![GitHub release](https://img.shields.io/github/release/rgglez/go-playground-uts-validator.svg)](https://github.com/rgglez/go-playground-uts-validator/releases/)
![GitHub stars](https://img.shields.io/github/stars/rgglez/go-playground-uts-validator?style=social)
![GitHub forks](https://img.shields.io/github/forks/rgglez/go-playground-uts-validator?style=social)

Custom validator for [github.com/go-playground/validator/v10](`github.com/go-playground/validator/v10`) that validates [UNIX timestamps](https://www.unixtimestamp.com/).

It supports `string`, integer types, and `*time.Time` (non-zero values).

## Installation

```bash
go get github.com/rgglez/go-playground-uts-validator
```

## Usage

Define a custom validator(s) using the fiber middleware:

```go
app.Use(
    gofibervalidator.New(gofibervalidator.Config{
        ContextKey: "validator",
        CustomValidations: map[string]validator.Func{
            "uts": isuts.ValidateUnixTimestamp,
        },
    }),
)
```

Use the `gofibervalidator.ValidateStruct` function to validate your input struct in the handler:

```go
	// Validate input
	defaults.Set(&input)
	if errors := gofibervalidator.ValidateStruct(c, input); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}
```

See the example in the [`examples`](examples) directory.

## Tests

Run the test suite with:

```bash
go test ./...
```

To include coverage:

```bash
go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out
```

### Test cases

| Group | Cases |
|-------|-------|
| `RegisterUTSValidator` | default tag, custom tag |
| String fields | valid numeric strings, empty string, non-numeric, float-like strings, strings with spaces |
| Signed integers | `int`, `int8`, `int16`, `int32`, `int64` |
| Unsigned integers | `uint`, `uint8`, `uint16`, `uint32`, `uint64` |
| `time.Time` | non-zero values (valid), zero value (invalid) |
| Pointer fields | `nil` pointer (invalid), non-nil pointer (valid) |
| Unsupported kinds | `float64`, `bool` |

## License

Copyright (C) 2026 Rodolfo González González.

Licensed under the Apache License, Version 2.0. Read the [LICENSE](LICENSE) file.
