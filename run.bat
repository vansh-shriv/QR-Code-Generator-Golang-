@echo off
echo QR Code Generator
echo =================
echo.
echo Choose mode:
echo 1. Web Interface (default)
echo 2. CLI Mode
echo.
set /p choice="Enter your choice (1 or 2): "

if "%choice%"=="2" (
    echo.
    set /p text="Enter text to encode: "
    set /p output="Enter output filename (default: qr-code.png): "
    if "%output%"=="" set output=qr-code.png
    go run main.go -cli -text "%text%" -output "%output%"
) else (
    echo.
    echo Starting web server on http://localhost:8080
    echo Press Ctrl+C to stop the server
    echo.
    go run main.go
) 