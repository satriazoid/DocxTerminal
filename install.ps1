# DocxTerminal Installer — PowerShell

Write-Host "📦 DocxTerminal Installer" -ForegroundColor Cyan
Write-Host ""

# Check Go
try {
    $goVersion = go version
    Write-Host "✓ Go found: $goVersion" -ForegroundColor Green
} catch {
    Write-Host "❌ Go not found. Install from https://go.dev/dl (1.24+)" -ForegroundColor Red
    exit 1
}

# Get GOPATH\bin
$goPath = (go env GOPATH)
$goBin = Join-Path $goPath "bin"

if (-not (Test-Path $goBin)) {
    New-Item -ItemType Directory -Path $goBin -Force | Out-Null
}
Write-Host "✓ Install dir: $goBin" -ForegroundColor Green

# Check PATH
$pathArray = $env:PATH -split ";"
if ($pathArray -notcontains $goBin) {
    Write-Host ""
    Write-Host "⚠️  $goBin not in PATH" -ForegroundColor Yellow
    Write-Host ""
    Write-Host "Add manually:" -ForegroundColor Yellow
    Write-Host "  Win+R > sysdm.cpl > Advanced > Environment Variables"
    Write-Host "  User Path > New > $goBin"
    Write-Host ""
}

# Build
Write-Host ""
Write-Host "🔨 Building dt..." -ForegroundColor Cyan
go build -o dt.exe .
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Build failed" -ForegroundColor Red
    exit 1
}

# Install
Write-Host "📂 Installing to $goBin\dt.exe..." -ForegroundColor Cyan
Move-Item -Path "dt.exe" -Destination (Join-Path $goBin "dt.exe") -Force

# Verify
$dtPath = Join-Path $goBin "dt.exe"
if (Test-Path $dtPath) {
    Write-Host ""
    Write-Host "✅ Install OK! Restart terminal, then try: dt" -ForegroundColor Green
} else {
    Write-Host ""
    Write-Host "⚠️  Installation may have failed" -ForegroundColor Yellow
}
