package react

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/dop251/goja"
)

var jsRender func(goja.FunctionCall) goja.Value

func initRenderer() {
	file, err := os.Open("client/public/bundle.js")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)

	v, err := vm.RunString(string(bytes))
	if err != nil {
		log.Fatal(err)
	}

	jsRender = v.Export().(map[string]interface{})["renderHtmlString"].(func(goja.FunctionCall) goja.Value)
}

// Render renders the react data to html string
func Render(url string, data interface{}) string {
	var m map[string]interface{}
	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &m)
	return jsRender(goja.FunctionCall{Arguments: []goja.Value{vm.ToValue(url), vm.ToValue(m)}}).String()
}
