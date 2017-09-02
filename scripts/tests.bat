@echo off

docker build -f .\Dockerfile-windows --tag buildkiteagent%BUILDKITE_BUILD_NUMBER% .
docker run --rm buildkiteagent%BUILDKITE_BUILD_NUMBER% go test ./...
docker rmi buildkiteagent%BUILDKITE_BUILD_NUMBER%