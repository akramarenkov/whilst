# Whilst

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/whilst.svg)](https://pkg.go.dev/github.com/akramarenkov/whilst)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/whilst)](https://goreportcard.com/report/github.com/akramarenkov/whilst)
[![Coverage Status](https://coveralls.io/repos/github/akramarenkov/whilst/badge.svg)](https://coveralls.io/github/akramarenkov/whilst)

## Purpose

Library that extends time.Duration from standard library with days, months and years

## Features

Without approximations

Allows for presence of spaces

Compatible with time.Duration

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
    whl, err := whilst.Parse("2y3mo10d23h59m58s10ms30µs10ns")
    if err != nil {
        panic(err)
    }

    fmt.Println(whl)
    fmt.Println(whl.When(time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC)))
    // Output:
    // 2y3mo10d23h59m58.01003001s
    // 2025-07-11 23:59:58.01003001 +0000 UTC
}
```
