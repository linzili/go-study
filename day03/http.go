package day03

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

func Http() {
	//get()
	//post()
	//client()
	//header()
	//server()
	forward()
}

func forward() {
	http.HandleFunc("/forward", func(writer http.ResponseWriter, request *http.Request) {
		director := func(request *http.Request) {
			request.URL.Scheme = "https"
			request.URL.Host = "golang.org"
			request.URL.Path = "/index"
		}

		proxy := httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(writer, request)
	})

	http.ListenAndServe(":8080", nil)
}

func server() {
	//http.Handle("/index", &MyHandler{})
	http.HandleFunc("/index", func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Println(responseWriter, "index")
	})
	http.ListenAndServe(":8080", nil)
}

type MyHandler struct {
}

func (h *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("my implement")

}

func get() {
	resp, err := http.Get("https://baidu.com")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	fmt.Println(string(content))
}

type Person struct {
	UserId   string
	Username string
	Age      int
	Address  string
}

func post() {
	person := Person{
		UserId:   "120",
		Username: "jack",
		Age:      18,
		Address:  "beijing",
	}

	json, _ := json.Marshal(person)
	fmt.Println(string(json))
	reader := bytes.NewReader(json)
	resp, err := http.Post("https://golang.org", "application/json;charset=utf-8", reader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	defer resp.Body.Close()
}

func client() {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "https://golang.org", nil)
	resp, _ := client.Do(request)
	content, _ := io.ReadAll(resp.Body)
	fmt.Println(string(content))
	defer resp.Body.Close()
}

func header() {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "https://golang.org", nil)
	request.Header.Add("Authorization", "123456")
	resp, _ := client.Do(request)
	fmt.Println(resp.Header)
	content, _ := io.ReadAll(resp.Body)
	fmt.Println(string(content))
	defer resp.Body.Close()
}
