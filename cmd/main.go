package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func init() {
	name, _ := os.LookupEnv("APP_NAME")
	level, _ := os.LookupEnv("LOG_LEVEL")
	fmt.Println("APP_NAME:", name)
	fmt.Println("LOG_LEVEL:", level)
}

func main() {
	http.HandleFunc("/api", rootHandler())

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalln(err)
	}
}

func rootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from API"))
	}
}
