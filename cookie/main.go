package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
)

const BASEURL = "http://localhost:18888"

func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	http.DefaultClient = &http.Client{
		Jar: jar,
	}
	u, _ := url.Parse(BASEURL)
	u.Path = "/cookie"
	for i := 0; i < 2; i++ {
		resp, err := http.Get(u.String())
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			panic(err)
		}

		log.Println(string(dump))
	}
}
