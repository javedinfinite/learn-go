package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
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

	// Lets see how to send post request similarly
	fmt.Println("****************Printing http post request*************************")
	
	// Simple post
	data := strings.NewReader("I am data to be sent {'test':'test_value'} ") // HTTP body must be a stream, converting string to stream
	resp2, err2 := http.Post(
		"https://example.com",
		"application/x-www-form-urlencoded",
		data,
	)
	if err2!=nil {
		panic(err2)
	}
	defer resp2.Body.Close()
	fmt.Println("printing post response", resp2.Status) // 405 as example.com doesn't accespt post request
	// Simple post end

	// Lets learn how we can send post with better approach for prod

	jsonData := []byte(`{"key1":"value1", "key2": "value2"}`) // everything mostly in slice of bytes as it is lowest common denominator of data, efficient and flexible, the backbone of Go’s I/O system
	req3, _ := http.NewRequest(
		"POST",
		"https://example.com",
		bytes.NewBuffer(jsonData), //  // HTTP body must be a stream, converting string to stream
	)
	req.Header.Set("Content-Type", "application/json")

	client2 := http.Client{}
	resp3, err3 := client2.Do(req3)

	if err3!=nil {
		panic(err3)
	}

	defer resp3.Body.Close()

	fmt.Println("Post request with client formation : ", resp3.Status) // 405 as example.com doesn't accespt post request

	// JSON: since most of the time data sent and received is JSON, in GO we constandtly need to encode (GO->JSON, json.Marshal) and decode (JSON->GO, json.Unmarshal)
	// For more check json.go file

	// Mostly we use json.NewDecoder(resp.Body).Decode(&user) and not unmarshal becasue that needs one more extra step of reading from stream first.
	// user is struct var
}