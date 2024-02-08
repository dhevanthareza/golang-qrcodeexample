package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	qrcode "github.com/skip2/go-qrcode"
	// "github.com/yeqown/go-qrcode/v2"
	// "github.com/yeqown/go-qrcode/writer/standard"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	domainBytes = "abcdefghijklmnopqrstuvwxyz0123456789"
)

func main() {
	runtime.GOMAXPROCS(8)
	// GENERATING URL
	urls := make([]string, 100)
	for i := range urls {
		urls[i] = generateRandomURL()
	}

	// GENERATING QRCODE
	start := time.Now()
	fmt.Println("Generating QRCODE with skip2/go-qrcode")
	// fmt.Println("Generating QRCODE with yeqown/go-qrcode")
	var wg sync.WaitGroup
	for index := range urls {
		wg.Add(1)
		// go yeqownqr(&index, &urls[index], &wg)
		go step2qr(&index, &urls[index], &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Elapsed time:", elapsed)

}

func step2qr(index *int, url *string, wg *sync.WaitGroup) {
	defer wg.Done()
	path := "generated_qr/" + strconv.Itoa(*index) + ".jpg"
	qrcode.WriteFile(*url, qrcode.Low, 256, path)

}

// func yeqownqr(index *int, url *string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	qrc, err := qrcode.New(*url)
// 	if err != nil {
// 		fmt.Printf("could not generate QRCode: %v", err)
// 		return
// 	}

// 	w, err := standard.New("./generated_qr/" + strconv.Itoa(*index) + ".jpeg")
// 	if err != nil {
// 		fmt.Printf("standard.New failed: %v", err)
// 		return
// 	}

// 	// save file
// 	if err = qrc.Save(w); err != nil {
// 		fmt.Printf("could not save image: %v", err)
// 	}

// }

func generateRandomURL() string {
	var sb strings.Builder
	sb.WriteString("http://")
	for i := 0; i < 10; i++ {
		sb.WriteByte(domainBytes[rand.Intn(len(domainBytes))])
	}
	sb.WriteByte('.')
	for i := 0; i < 3; i++ {
		sb.WriteByte(domainBytes[rand.Intn(len(domainBytes))])
	}
	sb.WriteByte('/')
	for i := 0; i < 10; i++ {
		sb.WriteByte(letterBytes[rand.Intn(len(letterBytes))])
	}
	return sb.String()
}
