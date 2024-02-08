package main

import (
	"strconv"

	go_qr "github.com/piglig/go-qr"
	// "image/color"
)

func processPiglig(index *int, url *string) {
	errCorLvl := go_qr.Low
	qr, err := go_qr.EncodeText(*url, errCorLvl)
	if err != nil {
		return
	}
	config := go_qr.NewQrCodeImgConfig(10, 4)
	err = qr.PNG(config, "generated_qr/"+strconv.Itoa(*index)+".png")
	if err != nil {
		return
	}
}
