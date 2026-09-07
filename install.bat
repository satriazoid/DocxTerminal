@echo off
REM DocxTerminal Installer — Windows

setlocal enabledelayedexpansion

echo 📦 DocxTerminal Installer
echo.

REM Check Go
where go >nul 2>nul
if errorlevel 1 (
    echo ❌ Go not found. Install from https://go.dev/dl (1.24+^)
    exit /b 1
)

for /f "tokens=3" %%i in ('go version') do set GO_VERSION=%%i
echo ✓ Go %GO_VERSION% found

REM Detect GOPATH\bin
for /f "delims=" %%i in ('go env GOPATH') do set GOPATH=%%i
set GOBIN=%GOPATH%\bin

if not exist "%GOBIN%" mkdir "%GOBIN%"
echo ✓ Install dir: %GOBIN%

REM Check PATH
echo %PATH% | find /i "%GOBIN%" >nul
if errorlevel 1 (
    echo.
    echo ⚠️  %GOBIN% not in PATH
    echo.
    echo Add manually:
    echo   Win+R ^> sysdm.cpl ^> Advanced ^> Environment Variables
    echo   User Path ^> New ^> %GOBIN%
    echo.
)

REM Build
echo.
echo 🔨 Building dt...
go build -o dt.exe .
if errorlevel 1 (
    echo ❌ Build failed
    exit /b 1
)

REM Install
echo 📂 Installing to %GOBIN%\dt.exe...
move /y dt.exe "%GOBIN%\dt.exe" >nul

REM Verify
where dt >nul 2>nul
if errorlevel 0 (
    echo.
    echo ✅ Install OK! Restart terminal, then try: dt
) else (
    echo.
    echo ⚠️  dt not in PATH yet. Restart terminal ^& add %GOBIN% to PATH first.
)

endlocal
