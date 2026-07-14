@echo off
REM Run the mental-health-id server from repository root (Windows CMD)
REM Usage: run.bat [args passed to `go run`]
cd /d "%~dp0"
rem forward any args to go run
go run ./cmd/mental-health-id %*
exit /b %ERRORLEVEL%