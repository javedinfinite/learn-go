package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func PrintHttpClient () {
	fmt.Println("*****************Printing http client****************************")

	// http is a protocol(rules) which is foundation of all web communication
	// Go gives a package: net/http  using which we can send requests to servers to fetch data or update/upload data. 
	resp, err := http.Get("https://gobyexample.com") // it's very shortcut way to send requesdt which hides layers like client and newRequest 

	if err!=nil {
		fmt.Println("http get request failed : ", err)
		return
	}

	// Think it like a finally block which we need to execute at any cost at the end of this function execution. 
	// Whether it is having error, succes etc to make sure the connection is closed to free network resource
	defer resp.Body.Close()

	// Check if the response is success to proceed further
	if resp.StatusCode !=200 {
		fmt.Println("Something went wrong : ", resp.Status)
		return
	}

	// resp.Body is not text, it is stream (water flowing through a pipe). io provides ReadAll() to read everything from stream and return []byte, error if wrong
	// io.ReadAll(resp.Body) reads entire http response into RAM as raw response, dangerous for big response.
	body, _ := io.ReadAll(resp.Body)

	fmt.Println("response body received : ", string(body)) // <html> response 

	// learning: when we are doing http.Get("https://example.com") req, _ := http.NewRequest("GET", "https://example.com", nil) client.Do(req)

	// Above is basic example. In prod that doesn't get used for some obvioud reasons so lets improve.
	// We need to build(for custom controlls like headers set) the request and then send the request
	// create http client so that we can configure the requests , add timeout, 
	// (NewRequest) → describes request (What)
	// (Client)     → sends request (How)
	// (Do)         → executes (Action)

	// Forming the request with definitions of method, url, headers, body etc..
	req, _ := http.NewRequest("GET", "https://gobyexample.com", nil)

	// Preparing how to send the request with timeout control, connection handling, transport behaviour
	client := &http.Client{
		Timeout: time.Second * 2,
	} // since it returns a struct, we need to use pointer to that struct so that we can modify it

	resp1, err1 := client.Do(req)
	if err1!=nil {
		fmt.Println("Get request failes with error", err1)
	}
	
	if resp1.StatusCode != http.StatusOK {
		fmt.Println("Something went wrong", resp1.Status)
		return
	}

	body1, _ := io.ReadAll(resp1.Body)
	fmt.Println("resp1 : ", string(body1[0]))

}