package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	domainBytes = "abcdefghijklmnopqrstuvwxyz0123456789"
)

func main() {
	// GENERATING URL
	urls := make([]string, 1000)
	for i := range urls {
		urls[i] = generateRandomURL()
	}

	// GENERATING QRCODE
	start := time.Now()
	// var wg sync.WaitGroup
	for index := range urls {
		// wg.Add(1)
		step2qr(&index, &urls[index])
		// processBoombuler(&index, &urls[index])
		// processPiglig(&index, &urls[index])
		// yeqownqr(&index, &urls[index])
	}
	// wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Elapsed time:", elapsed)

}

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
