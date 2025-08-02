package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func main() {
	// Parse command line flags
	cliMode := flag.Bool("cli", false, "Run in CLI mode")
	text := flag.String("text", "", "Text to encode in QR code (CLI mode)")
	output := flag.String("output", "qr-code.png", "Output file path (CLI mode)")
	port := flag.String("port", "8080", "Port for web server")
	flag.Parse()

	if *cliMode {
		runCLI(*text, *output)
	} else {
		runWebServer(*port)
	}
}

func runCLI(text, output string) {
	if text == "" {
		fmt.Println("Error: Please provide text to encode using -text flag")
		fmt.Println("Usage: go run main.go -cli -text \"Your text here\" -output qr-code.png")
		os.Exit(1)
	}

	// Create QR code
	err := qrcode.WriteFile(text, qrcode.Medium, 256, output)
	if err != nil {
		log.Fatal("Error generating QR code:", err)
	}

	fmt.Printf("QR code generated successfully: %s\n", output)
}

func runWebServer(port string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "QR Code Generator",
		})
	})

	r.POST("/generate", func(c *gin.Context) {
		text := c.PostForm("text")
		if text == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Text is required"})
			return
		}

		// Generate QR code
		filename := "static/qr-codes/generated-qr.png"
		err := qrcode.WriteFile(text, qrcode.Medium, 256, filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"filename": "/static/qr-codes/generated-qr.png",
		})
	})

	// Create static directories if they don't exist
	os.MkdirAll("static/qr-codes", 0755)

	fmt.Printf("Web server running on http://localhost:%s\n", port)
	r.Run(":" + port)
} 