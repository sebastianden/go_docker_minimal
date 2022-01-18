package main

import (
	"net/http";
	"io"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
		<head>
		<title>Hello World &#10084;&#65039;</title>
		</head>
		<body>
		Hello World &#10084;&#65039;!
		</body>
		</html>`,
	)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
