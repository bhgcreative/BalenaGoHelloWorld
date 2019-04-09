package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello,Brian Harlon Green!!!!!!!)\nI would like to welcome you to your portal."))
	})
	addr := ":80"
	fmt.Println("Example app listening on port ", addr)
	fmt.Println("Hello, Brian Harlon Green ")
	log.Fatal(http.ListenAndServe(addr, nil))
}
