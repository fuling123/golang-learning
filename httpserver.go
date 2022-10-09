package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/*
接收客户端 request，并将 request 中带的 header 写入 response header
读取当前系统的环境变量中的 VERSION 配置，并写入 response header
Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
当访问 localhost/healthz 时，应返回 200
*/
func main() {
	http.HandleFunc("/getHeader", getHeader)
	http.HandleFunc("/getVersion", getVersion)
	http.HandleFunc("/getLog", getLog)
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getHeader(response http.ResponseWriter, request *http.Request) {
	headers := request.Header
	for header := range headers {
		values := headers[header]
		for index, _ := range values {
			values[index] = strings.TrimSpace(values[index])

		}
		println(header + "=" + strings.Join(values, ","))
		response.Header().Set(header, strings.Join(values, ","))
	}
	fmt.Fprintln(response, "Header:", headers)
	io.WriteString(response, "success")

}

func getVersion(response http.ResponseWriter, request *http.Request) {
	envStr := os.Getenv("windir")
	response.Header().Set("windir", envStr)
	io.WriteString(response, "success")
}

func getLog(response http.ResponseWriter, request *http.Request) {
	form := request.RemoteAddr
	println("Client->ip:port=" + form)
	ipStr := strings.Split(form, ":")
	println("Client->ip=" + ipStr[0])
	println("Client->response code=" + strconv.Itoa(http.StatusOK))
	io.WriteString(response, "success")

}

func healthz(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	io.WriteString(response, "success")
}
