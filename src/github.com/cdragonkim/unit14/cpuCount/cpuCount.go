package cpuCount

import (
	"fmt"
	"runtime"
)

func CPUcnt() {
	fmt.Println("CPU Count : ", runtime.NumCPU())
}
