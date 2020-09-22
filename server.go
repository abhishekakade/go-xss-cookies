package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	keys, present := r.URL.Query()["key"]
	if !present || len(keys[0]) < 1 {
		log.Println("No key found in the URL!")
		return
	} else {
		key := []byte(keys[0])
		error := ioutil.WriteFile("keys.txt", key, 0644)

		if error != nil {
			log.Fatal(error)
		}
		log.Println("Key successfully written to the file.")
		return
	}
}
