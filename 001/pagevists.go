// track the number of page visits w/ cookies

package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.io", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("my-cookie-counter")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie-counter",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}

	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(res, cookie)
	io.WriteString(res, cookie.Value)
}
