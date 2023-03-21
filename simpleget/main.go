package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const URL = "http://localhost:18888"

func get() {
	values := url.Values{
		"query": []string{"hello world"},
	}
	resp, err := http.Get(URL + "?" + values.Encode())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))

	// 文字列で "200 OK"
	log.Println("Status:", resp.Status)
	// 数値で 200
	log.Println("StatusCode:", resp.StatusCode)
	// ヘッダーの表示
	log.Println("Headers:", resp.Header)
	// Content-Lengthを表示
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
}

func post() {
	file, err := os.Open("simpleget/main.go")
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(URL, "text/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}

func postByte() {
	reader := strings.NewReader("テキスト")
	resp, err := http.Post(URL, "text/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}

func postForm() {
	values := url.Values{
		"test": []string{"value"},
	}
	resp, err := http.PostForm(URL, values)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}

func postMultipart() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Gopher")
	fileWriter, err := writer.CreateFormFile("thumbnail", "data/gopher.jpg")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("data/gopher.jpg")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post(URL, writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}

func main() {
	get()
	postForm()
	post()
	postByte()
	postMultipart()
}
