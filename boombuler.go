package main

import (
	"image/png"
	"os"
	"strconv"

	"github.com/boombuler/barcode"
	qr "github.com/boombuler/barcode/qr"
)

func processBoombuler(index *int, url *string) {
	qrCode, _ := qr.Encode(*url, qr.L, qr.Encoding(qr.L))

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create("generated_qr/" + strconv.Itoa(*index) + ".png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)

}
