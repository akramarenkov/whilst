# Whilst

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/whilst.svg)](https://pkg.go.dev/github.com/akramarenkov/whilst)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/whilst)](https://goreportcard.com/report/github.com/akramarenkov/whilst)
[![Coverage Status](https://coveralls.io/repos/github/akramarenkov/whilst/badge.svg)](https://coveralls.io/github/akramarenkov/whilst)

## Purpose

Library that extends time.Duration from standard library with days, months and years

## Features

Without approximations

Allows for presence of spaces in a string representation

Compatible with time.Duration in terms of parsing and conversion to string

## Usage

Example:

```go
package main

import (
    "fmt"
    "time"

    "github.com/akramarenkov/whilst"
)

func main() {
    from := time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC)

    whl, err := whilst.Parse("2y")
    if err != nil {
        panic(err)
    }

    fmt.Println(whl)
    fmt.Println(whl.When(from))
    fmt.Println(whl.Duration(from))
    // Output:
    // 2y
    // 2025-04-01 00:00:00 +0000 UTC
    // 17544h0m0s
}
```
