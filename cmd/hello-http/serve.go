package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(
		writer http.ResponseWriter,
		request *http.Request) {
			// fmt.Fprintln(writer, "<h1>Hello Serve!</h1>")
			fmt.Fprintf(writer, "<h1>Hello Serve %s!</h1>", request.FormValue("name"))
	})

	http.ListenAndServe(":8080", nil)
}
