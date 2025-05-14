package server

import (
	"fmt"
	"log"
	"net/http"
)

func Start(port string) {
	fmt.Println("starting server on port: ", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go server on port:", port)
	})

	addr := ":" + port
	fmt.Println("Server running at http://localhost" + addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
