package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/stores/number/", singleHandler)
	http.HandleFunc("/stores/all", allHandler)
	http.ListenAndServe(":8080", nil)
}

func allHandler(w http.ResponseWriter, r *http.Request) {
	res, err := getall()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	data, err := prepJson(res)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", string(data))
}

func singleHandler(w http.ResponseWriter, r *http.Request) {
	sn := r.URL.Path[len("/stores/number/"):]
	if len(sn) == 0 {
		fmt.Fprintf(w, "%s", "Invalid Store Number")
		return
	}

	res, err := getstore(sn)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	data, err := prepJson(res)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", string(data))
}

func prepJson(res Response) ([]byte, error) {
	data, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	data = bytes.Replace(data, []byte("\\u003c"), []byte("<"), -1)
	data = bytes.Replace(data, []byte("\\u003e"), []byte(">"), -1)
	data = bytes.Replace(data, []byte("\\u0026"), []byte("&"), -1)

	return data, nil
}
