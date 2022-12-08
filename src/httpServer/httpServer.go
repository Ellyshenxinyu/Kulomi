package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"runtime"
	"strconv"
)

/*
编写一个HTTP服务器
1、接受客户端request,并将request中带入的header写入response header
2、读取当前系统的环境变量中的VERSION配置,并写入response header
3、Server端记录访问日志包括客户端IP,HTTP返回码,输出到server端的标准输出
4、当访问localhost/healthz时,应返回200
*/
/*
 log.Print() 函数在控制台屏幕上打印带有时间戳的指定消息
 log.Fatal() 打印输出内容；退出应用程序；defer不执行
 Header类型 map[string][]string
 strconv.Itoa() int=>string
 strconv.Atoi() string=>int
 设置http的头，状态码，body的顺序如下：w.Header().Set() => w.WriteHeader() => w.Write()
 https://blog.csdn.net/LngZd/article/details/115359169
*/
const ver = "GOVERSION"

func main() {
	http.HandleFunc("/", httpHandler)
	http.HandleFunc("/healthz", healthz)

	// 默认listen 0.0.0.0:80
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "Hello!")
	if r.Header == nil || len(r.Header) == 0 {
		log.Println("header为空")
	} else {
		i := 0
		//#1、将request中带入的header写入response header
		for key, values := range r.Header {
			for _, value := range values {
				w.Header().Set(key, value)
				i++
				log.Printf("%d : %s = %s\n", i, key, value)
			}
		}
		fmt.Println()
	}

	//#2、读取当前系统的环境变量中的VERSION配置,并写入response header
	// version := os.Getenv(ver) 获取不到
	version := runtime.Version()
	w.Header().Set(ver, version)
	log.Printf("%s = %s\n", ver, version)

	//#3、Server端记录访问日志包括客户端IP,HTTP返回码,输出到server端的标准输出
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("err:", err)
	}
	// ::1是IPv6中的环回地址，把它看作是127.0.0.1的IPv6版本
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	log.Printf("%s = %s", "IP", ip)

	//res, _ := http.Get("")
	//println(res.StatusCode)
	w.WriteHeader(http.StatusInternalServerError)
	log.Println("statusCode:", http.StatusInternalServerError)

	// 最后输出 不然状态码设置无效！
	io.WriteString(w, "Hello!")
}

// #4、访问localhost:8080/healthz时,返回200
func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, strconv.Itoa(http.StatusOK))
}
