package main

import (
	flag "github.com/ogier/pflag"
	"fmt"
	"os"
	uri "net/url"
	"net/http"
	"io/ioutil"
	"strings"
)

var (
	url  string
)

func showHelp() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		showHelp()
	}

	if  os.Args[1] == "-s" && checkValidity(url){

		data := uri.Values{}
		data.Set("url", url)

		client := &http.Client{}
		req, _ := http.NewRequest("POST", "https://fujitora.herokuapp.com/api/v1/shorten", strings.NewReader(data.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")


		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	} else if os.Args[1] == "-e" && checkValidity(url){
		//TO_DO
		fmt.Print("I don sabi the work")
	} else {
		showHelp()
	}

}

func isValidUrl(url string) bool {
	_, err := uri.ParseRequestURI(url)
	if err != nil {
		return false
	} else {
		return true
	}
}

func checkValidity(url string) bool {
	var valid bool
	if isValidUrl(url) {
		valid = true
	} else {
		fmt.Print("Please input a valid url \n")
		valid = false
		os.Exit(1)
	}
	return valid
}

func init() {
	flag.StringVarP(&url, "shrink", "s", "", "shortens a url")
	flag.StringVarP(&url, "elongate", "e", "", "Returns the original url")
}