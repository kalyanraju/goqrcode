package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Define a POST endpoint for generating QR codes
	r.POST("/generate", func(c *gin.Context) {
		var request struct {
			Content string `json:"content" binding:"required"`
			Size    int    `json:"size" binding:"required"`
		}

		// Bind JSON input
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Generate the QR code
		png, err := qrcode.Encode(request.Content, qrcode.Medium, request.Size)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
			return
		}

		// Respond with the QR code as a PNG image
		c.Data(http.StatusOK, "image/png", png)
	})

	// Start the server
	r.Run(":8080") // Listen on port 8080
}
