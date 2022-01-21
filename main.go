package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+basicAuth("username1", "password123"))
	return nil
}

func main() {
	url := "https://icanhazdadjoke.com/"

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	resp, err := client.Get(url)
	fmt.Println("\nAPI Joke1: ", resp)

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Dadjoke CLI (github.com/yonathanavila)")

	resp, err = client.Do(req)

	if err != nil {
		// handle error
		log.Println("Ocurrio un error: ", err)
	}
	fmt.Println("\nAPI Joke2: ", resp)

	defer resp.Body.Close()

	fmt.Println("\nAPI Joke3: ", resp)
}
