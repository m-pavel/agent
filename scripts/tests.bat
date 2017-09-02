@echo off

echo '+++ Running tests'
docker run --rm \
  -v %cd%:/go/src/github.com/buildkite/agent \
  -W /go/src/github.com/buildkite/agent \
  golang:1.9-windowsservercore \
  go test ./...
