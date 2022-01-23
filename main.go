package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	urlGet := "https://icanhazdadjoke.com/"
	urlPost := "https://httpbin.org/post"
	values := map[string]string{"name": "John Doe", "occupation": "gardener"}

	fmt.Println(sendGet(urlGet))
	sendPost(values, urlPost)

}

func sendGet(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
	return string(body)
}

func sendPost(data map[string]string, url string) {

	json_data, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])

}
