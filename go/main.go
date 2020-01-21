package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/fatal", func(w http.ResponseWriter, r *http.Request) {
		a := 100
		b := 0
		x := a / b
		s := strconv.Itoa(x)

		fmt.Fprintf(w, s)
	})

	http.HandleFunc("/err", func(w http.ResponseWriter, r *http.Request) {
		log.Println("NG 400")
		//log.Fatalln("NG 400") >> exit status 1
		//log.Panicln("NG 400") >> panic

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "NG 400")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("OK 200")

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK 200")
	})

	http.ListenAndServe(":3000", nil)
}
