package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const error = "JU588a9a9b4bf35c1b6f0d39530f: 2016-12-27T00:37:32.847886+00:00@J5ROperation failed"
const success = "JY588a9a9b4bf35c1b6f0d39530f: 2016-12-27T00:37:32.847886+00:00J0ROperation SuccessfulX"

var count = 10

type increaseBody struct {
	Value int `json:"value"`
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/", handler)
	http.HandleFunc("/increase", increaseTimeSleep)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	timeSleep := time.Duration(time.Second * time.Duration(count))
	time.Sleep(timeSleep)
	fmt.Fprintf(responseWriter, zeroDollar())

}

func increaseTimeSleep(responseWriter http.ResponseWriter, request *http.Request) {
	body := increaseBody{}
	defer request.Body.Close()
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}

	json.Unmarshal(requestBody, &body)

	count += body.Value
}

func health(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "ok")
}

func zeroDollar() string {
	return success
}
