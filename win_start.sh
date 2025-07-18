#!/bin/sh
set -e # Exit early if any commands fail

(
    cd "$(dirname "$0")" # Ensure compile steps are run within the repository directory
    GOOS=windows GOARCH=amd64 go build -o ./build/win/build-http-server-go.exe app/*.go
)

exec ./build/win/build-http-server-go.exe "$@"