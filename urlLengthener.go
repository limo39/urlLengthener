package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	reader := bufio.NewReader(os.Stdin)


	fmt.Println("URL Lengthener")
	fmt.Println("--------------")
	fmt.Print("Enter the URL to lengthen: ")
	originalURL, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}


	originalURL = strings.TrimSpace(originalURL)

	if !strings.HasPrefix(originalURL, "http") {
		fmt.Println("Warning: This doesn't look like a complete URL (missing http/https)")
	}


	lengthenedURL := lengthenURL(originalURL)
	fmt.Println("\nOriginal URL:", originalURL)
	fmt.Println("Lengthened URL:", lengthenedURL)


	fmt.Print("\nPress enter to exit...")
	reader.ReadString('\n')
}

func lengthenURL(original string) string {
	u, err := url.Parse(original)
	if err != nil {
		return original
	}


	query := u.Query()
	

	query.Add("utm_source", randomString(8))
	query.Add("utm_medium", randomString(5))
	query.Add("utm_campaign", randomString(10))
	query.Add("ref", randomString(12))
	query.Add("fbclid", randomString(15))
	

	for i := 0; i < 5; i++ {
		paramName := randomString(rand.Intn(5) + 3)
		paramValue := randomString(rand.Intn(10) + 5)
		query.Add(paramName, paramValue)
	}


	u.RawQuery = query.Encode()


	if rand.Intn(2) == 1 {
		u.Fragment = randomString(8)
	}

	return u.String()
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
