package server

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/dop251/goja"
)

func (s *server) initRenderFunc() {
	file, err := os.Open("client/public/built/bundle.js")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)

	s.jsVM = goja.New()
	v, err := s.jsVM.RunString(string(bytes))
	if err != nil {
		log.Fatal(err)
	}

	s.jsRenderFunc = v.Export().(map[string]interface{})["genHtmlString"].(func(goja.FunctionCall) goja.Value)
}
