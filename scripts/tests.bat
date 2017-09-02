@echo off

echo '+++ Running tests'
docker run --rm ^
  -v %cd%:c:\gopath\src\github.com\buildkite\agent ^
  -w c:\go\src\github.com\buildkite\agent ^
  golang:1.9-nanoserver ^
  go test ./...
