package main

import (
	"strings"
	"./goproxy"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	myProxy := goproxy.NewProxyHttpServer()

	log.Println("Serving")

	myProxy.OnResponse().DoFunc(func(sitesResponse *http.Response, context *goproxy.ProxyCtx)*http.Response{
		body, _ := ioutil.ReadAll(sitesResponse.Body)
			if (strings.Contains(string(body), "new RTCPeerConnection")){
				println ("contains")
			}

		log.Println("response findQuiz==>", string(body))
		return sitesResponse
	})

	log.Fatal(http.ListenAndServe(":8080", myProxy))

}

