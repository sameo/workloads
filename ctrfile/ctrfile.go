package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	var loop = flag.Int("loop", 100, "help message for flagname")
	var fileName = flag.String("file", "/bin/zfgrep", "help message for flagname")

	flag.Parse()

	buffer := make([]byte, 512)
	
	for i:= 0; i < *loop; i++ {
		f, err := os.Open(*fileName)
		if err != nil {
			log.Fatal(err)
		}

		_, err = f.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}

		f.Close()
	}
}
