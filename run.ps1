#!/usr/bin/env pwsh

Write-Host "QR Code Generator" -ForegroundColor Cyan
Write-Host "=================" -ForegroundColor Cyan
Write-Host ""

Write-Host "Choose mode:" -ForegroundColor Yellow
Write-Host "1. Web Interface (default)" -ForegroundColor White
Write-Host "2. CLI Mode" -ForegroundColor White
Write-Host ""

$choice = Read-Host "Enter your choice (1 or 2)"

if ($choice -eq "2") {
    Write-Host ""
    $text = Read-Host "Enter text to encode"
    $output = Read-Host "Enter output filename (default: qr-code.png)"
    
    if ([string]::IsNullOrEmpty($output)) {
        $output = "qr-code.png"
    }
    
    try {
        go run main.go -cli -text $text -output $output
        if ($LASTEXITCODE -eq 0) {
            Write-Host "QR code generated successfully!" -ForegroundColor Green
        }
    }
    catch {
        Write-Host "Error: $_" -ForegroundColor Red
    }
} else {
    Write-Host ""
    Write-Host "Starting web server on http://localhost:8080" -ForegroundColor Green
    Write-Host "Press Ctrl+C to stop the server" -ForegroundColor Yellow
    Write-Host ""
    
    try {
        go run main.go
    }
    catch {
        Write-Host "Error: $_" -ForegroundColor Red
    }
} 