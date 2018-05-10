package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var loop = flag.Int("loop", 100, "help message for flagname")

	flag.Parse()

	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "gvisor")
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 128)
	tmpfileName := tmpfile.Name()
	tmpfileLen := len(content)
	
	defer os.Remove(tmpfileName) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	
	for i:= 0; i < *loop; i++ {
		f, err := os.Open(tmpfileName)
		if err != nil {
			log.Fatal(err)
		}

		nBytes, err := f.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}

		if nBytes != tmpfileLen {
			log.Fatal(fmt.Errorf("read %d expected %d\n", nBytes, tmpfileLen))
		}

		f.Close()
	}
}
