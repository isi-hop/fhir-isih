package instmemory

import (
	"fmt"
	"runtime"

	"github.com/mackerelio/go-osstat/memory"
)

func MemoryShow() {
	memory, err := memory.Get()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	// tester version du système
	//doit être une ubuntu version >= 20.04 LTS
	fmt.Println("Architecture : " + runtime.GOARCH + "_" + runtime.GOOS)
	fmt.Printf("Memoire Totale %d Go \n", memory.Total/1000/1024/1024)
	fmt.Printf("Memoire Libre %d Go \n", memory.Free/1000/1024/1024)

}

func Arch_test(name string) bool {
	return runtime.GOOS == name
}
