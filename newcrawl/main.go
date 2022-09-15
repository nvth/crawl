package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/steelx/extractlinks"
)

var (
	// ignore SSL
	// http transport

	// config of transport
	config = &tls.Config{
		// skip ssl cert
		InsecureSkipVerify: true,
	}
	// TLS config
	transport = &http.Transport{
		TLSClientConfig: config,
	}
	netClient = &http.Client{
		Transport: transport,
	}

	queue = make(chan string)

	hasVisited = make(map[string]bool)
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		fmt.Println("Missing URL")
		os.Exit(1)
	}

	baseURL := arguments[0]
	go func() {
		queue <- arguments[0]
	}()

	for href := range queue {
		if !hasVisited[href] {
			crawlURL(href)
		}
	}
	fmt.Println("baseURL", baseURL)

	crawlURL(baseURL)
}

func crawlURL(href string) {

	hasVisited[href] = true
	fmt.Printf("Crawling url -> %v \n", href)
	response, err := netClient.Get(href)
	checkErr(err)
	defer response.Body.Close()

	links, err := extractlinks.All(response.Body)
	checkErr(err)

	for _, link := range links {
		absolubteURL := toFixedURL(link.Href, href)
		go func() {
			queue <- absolubteURL
		}()
	}
}

func toFixedURL(href, baseURL string) string {
	uri, err := url.Parse(href)

	if err != nil {
		return ""
	}
	base, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}

	toFixedUri := base.ResolveReference(uri)
	// host from base
	// path from uri
	return toFixedUri.String()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
