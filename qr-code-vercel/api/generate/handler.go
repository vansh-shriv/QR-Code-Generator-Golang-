package main

import (
	"encoding/json"
	"github.com/skip2/go-qrcode"
	"net/http"
	"os"
)

type Request struct {
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var req Request
		json.NewDecoder(r.Body).Decode(&req)

		if req.Text == "" {
			http.Error(w, "Missing text field", http.StatusBadRequest)
			return
		}

		qr, err := qrcode.Encode(req.Text, qrcode.Medium, 256)
		if err != nil {
			http.Error(w, "QR generation failed", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		w.Write(qr)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}

