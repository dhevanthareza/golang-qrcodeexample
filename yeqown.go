package main

import (
	"fmt"
	"strconv"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func yeqownqr(index *int, url *string) {
	// defer wg.Done()
	qrc, err := qrcode.New(*url)
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
		return
	}

	w, err := standard.New("./generated_qr/" + strconv.Itoa(*index) + ".jpeg")
	if err != nil {
		fmt.Printf("standard.New failed: %v", err)
		return
	}

	// save file
	if err = qrc.Save(w); err != nil {
		fmt.Printf("could not save image: %v", err)
	}

}
