package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func amIWorking() {
	neo := make(chan string)
	neo <- "one"
	value := <-neo
	print(value)
}

func amIAStuckWorkerFunc() {
	for {
		time.Sleep(10 * time.Second)
		fmt.Println("Alive")
	}
}

func httpCall(path string) {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}
	httpRequest, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		fmt.Println("1")
		panic(err)
	}

	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		fmt.Println("2")
		panic(err)
	}

	if httpResponse.StatusCode != http.StatusOK {
		panic("Not a 200 OK")
	}
	defer httpResponse.Body.Close()

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func main() {
	// 1 - Is this working? Why? - Fix me.
	amIWorking()
	return // Remove me once question answered

	// 2 - Ran like this, what's happening ? - Fix me.
	go amIAStuckWorkerFunc()
	return // Remove me once question answered

	// 3 - Find a way to cancel me mid-run when desired.
	go amIAStuckWorkerFunc()
	return // Remove me once question answered

	// 4 - Fix this routine, then use it to send multiple HTTP requests and retrieve the results
	httpUrls := []string{"https://www.google.com", "https://www.google.c", "https://www.bing.com", "https://www.github.com"}
	go httpCall(httpUrls[0])
	return // Remove me once question answered
}
