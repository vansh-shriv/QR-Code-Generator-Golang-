# QR Code Generator

A simple and elegant QR code generator built with Go, featuring both CLI and web interfaces.

## Features

- **CLI Mode**: Generate QR codes from command line
- **Web Interface**: Beautiful, responsive web UI for generating QR codes
- **Download Support**: Download generated QR codes as PNG files
- **Modern Design**: Clean and intuitive user interface
- **Cross-platform**: Works on Windows, macOS, and Linux

## Installation

1. Make sure you have Go 1.21 or later installed
2. Clone or download this repository
3. Navigate to the project directory
4. Install dependencies:

```bash
go mod tidy
```

## Usage

### Web Interface (Default)

To start the web server:

```bash
go run main.go
```

The web interface will be available at `http://localhost:8080`

You can also specify a custom port:

```bash
go run main.go -port 3000
```

### CLI Mode

To generate a QR code from the command line:

```bash
go run main.go -cli -text "Your text here" -output qr-code.png
```

#### CLI Options

- `-cli`: Enable CLI mode
- `-text`: Text or URL to encode in the QR code (required)
- `-output`: Output file path (default: qr-code.png)
- `-port`: Port for web server (default: 8080)

#### CLI Examples

```bash
# Generate QR code for a URL
go run main.go -cli -text "https://example.com" -output website-qr.png

# Generate QR code for plain text
go run main.go -cli -text "Hello, World!" -output hello-qr.png

# Generate QR code for contact information
go run main.go -cli -text "BEGIN:VCARD\nVERSION:3.0\nFN:John Doe\nTEL:+1234567890\nEND:VCARD" -output contact-qr.png
```

## Building

To build the executable:

```bash
go build -o qr-generator main.go
```

Then you can run it:

```bash
# Web mode
./qr-generator

# CLI mode
./qr-generator -cli -text "Your text" -output qr.png
```

## Project Structure

```
qr-code-generator/
├── main.go              # Main application file
├── go.mod               # Go module file
├── templates/
│   └── index.html       # Web interface template
├── static/
│   └── qr-codes/        # Generated QR codes (created automatically)
└── README.md           # This file
```

## Dependencies

- [gin-gonic/gin](https://github.com/gin-gonic/gin): Web framework
- [skip2/go-qrcode](https://github.com/skip2/go-qrcode): QR code generation library

## Features

### Web Interface
- Beautiful, responsive design
- Real-time QR code generation
- Download functionality
- Error handling and validation
- Mobile-friendly interface

### CLI Interface
- Fast command-line QR code generation
- Customizable output file names
- Support for various text types (URLs, plain text, contact info, etc.)

## QR Code Types Supported

The application can generate QR codes for:
- URLs and web addresses
- Plain text messages
- Contact information (vCard format)
- Email addresses
- Phone numbers
- WiFi network credentials
- And more!

## License

This project is open source and available under the MIT License. 