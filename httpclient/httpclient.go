package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var loop = flag.Int("loop", 100, "help message for flagname")
	var url = flag.String("url", "http://172.17.0.1:8080/index.html", "help message for flagname")

	flag.Parse()

	for i:= 0; i < *loop; i++ {
		res, err := http.Get(*url)
		if err != nil {
			log.Fatal(err)
		}
		_, err = ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}
