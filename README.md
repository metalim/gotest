# Go test action

[![Tests workflow](https://github.com/metalim/gotest/actions/workflows/tests.yml/badge.svg)](https://github.com/metalim/gotest/actions/workflows/tests.yml)

tests fail due to YAML issue with numbers: `go-version: 1.20` is treated as `go-version: 1.2`. The fix is to use `go-version: "1.20"`
