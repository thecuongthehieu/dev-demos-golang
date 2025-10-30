#!/bin/sh

# brew install lp_solve

# go mod tidy

export CGO_CFLAGS="-I/opt/homebrew/opt/lp_solve/include"
export CGO_LDFLAGS="-L/opt/homebrew/opt/lp_solve/lib"

go run ./packages/libstest/golp/main.go