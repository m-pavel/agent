@echo off

docker build -f .\Dockerfile-windows --tag buildkiteagent%BUILDKITE_BUILD_NUMBER% .
docker run --rm -it buildkiteagent%BUILDKITE_BUILD_NUMBER% go test ./...
docker rm buildkiteagent%BUILDKITE_BUILD_NUMBER%