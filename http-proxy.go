package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	esEndpointStr := os.Getenv("ES_ENDPOINT")
	if esEndpointStr == ""{
		panic("ES_ENDPOINT env not found")
	}
	esEndpoints := strings.Split(esEndpointStr, ",")

	index := 0
	//流量转发后端策略
	GetEndpoint := func() string {
		if len(esEndpoints) == 0 {
			return ""
		}
		index = (index + 1) % len(esEndpoints)
		return esEndpoints[index]
	}

	//处理请求
	client := &http.Client{}
	actionFunc := func(w http.ResponseWriter, r *http.Request, action string) {
		ip := GetEndpoint()
		if ip == "" {
			w.Write([]byte("no backend"))
			return
		}
		fmt.Println(ip+r.URL.RequestURI())
		req, err := http.NewRequest(r.Method, "http://"+ip+r.URL.RequestURI(), r.Body)
		if err != nil {
			return
		}
		req.Header.Set("Content-type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		b, _ := ioutil.ReadAll(resp.Body)
		w.Write(b)
		fmt.Println("receive proxy from ", r.RemoteAddr, " redirect to ", ip)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		actionFunc(w, r, "hello")
	})

	fmt.Println("esEndpointStr: ", esEndpointStr)
	fmt.Println("ListenAndServe: 0.0.0.0:8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
