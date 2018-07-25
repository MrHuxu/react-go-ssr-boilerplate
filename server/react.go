package server

import (
	"io/ioutil"
	"log"
	"os"
)

func (s *server) initReact() {
	file, err := os.Open("client/public/built/bundle.js")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	s.react = string(bytes)
}
