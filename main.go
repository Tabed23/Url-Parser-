package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/steelx/extractlinks"
)

var (
	config = &tls.Config{
		InsecureSkipVerify: true,
	}
	transpot = &http.Transport{
		TLSClientConfig: config,
	}
	netClinet = &http.Client{
		Transport: transpot,
	}
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("enter the url which you want to parse\n")
		os.Exit(0)
	}
	url := args[0]
	response, err := netClinet.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	links := extractLinks(response)
	writeOnFile(links)
	printLinks(links)
}
func extractLinks(res *http.Response) []extractlinks.Link {
	data, err := extractlinks.All(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
func printLinks(link []extractlinks.Link) {
	fmt.Println(" All the Links")
	for _, links := range link {
		fmt.Printf("link %v \n", links)
	}
}

func writeOnFile(links []extractlinks.Link) {
	file, err := os.Create("links.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, link := range links {
		file.WriteString(string(link.Href) + "\t" + string(link.Text) + "\n")
	}
}
