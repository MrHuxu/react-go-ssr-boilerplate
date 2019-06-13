package react

import (
	"github.com/dop251/goja"
)

// VM exports the instance of js runtime
var vm *goja.Runtime

func initVM() {
	vm = goja.New()
}
