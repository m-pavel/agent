@echo off

set image_name=buildkite-agent-%BUILDKITE-BUILD_NUMBER%

docker build -f .\Dockerfile-windows --tag %image_name% .
docker run --rm -it %image_name% go test ./...
docker rm %image_name%