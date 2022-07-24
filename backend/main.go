package main

import (
	"backend/helper"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helper.IndexHandler)
	http.Handle("/asset/", http.StripPrefix("../frontend/identik2k22/asset/", http.FileServer(http.Dir("./asset"))))
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe("localhost:9000", nil)

}
