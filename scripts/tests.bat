@echo off

docker build -f .\Dockerfile-windows --tag buildkiteagent%BUILDKITE_BUILD_NUMBER% . || goto :error
docker run --rm buildkiteagent%BUILDKITE_BUILD_NUMBER% go test ./... || goto :error
goto :EOF

:error
set previous_errorlevel=%errorlevel%
docker rmi buildkiteagent%BUILDKITE_BUILD_NUMBER%
echo Failed with error #%previous_errorlevel%.
exit /b %previous_errorlevel%