package main

import ("fmt"
		"net/url")

func PrintUrlParsing () {

	fmt.Println("*******************Printing UrlParsing***********************")
	// Go gives you a solid standard library package for this: net/url. Using the library we can segregate the different parts of the url for our use case.

	rawURL := "https://example.com:8080/search?q=golang&q2=test#section1"

	parsedUrl, err := url.Parse(rawURL)
	if err!=nil {
		panic(err)
	}
	fmt.Println("Parsed URL : ", parsedUrl)

	fmt.Println("Scheme", parsedUrl.Scheme)
	fmt.Println("Host", parsedUrl.Host)
	fmt.Println("HostName", parsedUrl.Hostname())
	fmt.Println("Path", parsedUrl.Path)
	fmt.Println("Rawquery", parsedUrl.RawQuery)
	fmt.Println("query value:", parsedUrl.Query().Get("q"))
	fmt.Println("query value q2:", parsedUrl.Query().Get("q2"))
	fmt.Println("Fragment", parsedUrl.Fragment)
	fmt.Println("port", parsedUrl.Port())

	t := parsedUrl.Query()
	t.Set("q2", "sdd")
	parsedUrl.RawQuery = t.Encode()
	fmt.Println("update query value q2:", parsedUrl)

}