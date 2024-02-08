package main

import (
	"strconv"

	"github.com/skip2/go-qrcode"
)

func step2qr(index *int, url *string) {
	// defer wg.Done()
	path := "generated_qr/" + strconv.Itoa(*index) + ".jpg"
	qrcode.WriteFile(*url, qrcode.Low, 256, path)

}
