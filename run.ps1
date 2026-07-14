# Run the mental-health-id server from repository root (PowerShell)
# Usage: .\run.ps1 <args>
Push-Location $PSScriptRoot
try {
    & go run ./cmd/mental-health-id @args
} finally {
    Pop-Location
}