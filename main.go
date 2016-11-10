package main

import (
	"os"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"io/ioutil"
)

func main() {
	pidOfProcs := getAllProcessPids()
	memOfProcs := make([]int, len(pidOfProcs))
	for i, pid := range pidOfProcs {
		memOfProcs[i] = getUsedMemory(pid)
	}
	big := biggestProcess(memOfProcs)
	if big > 0 {
		killProcess(pidOfProcs[big])
	}
}

func getAllProcessPids() []string {
	processes, err := filepath.Glob("/proc/[0-9]*") //get all files in /proc
	if err != nil {
		fmt.Println(err)
	}
	return processes
}

func getUsedMemory(pid string) int {
	status, err := ioutil.ReadFile(pid+"/statm")
	if err != nil {
		fmt.Println(err)
	}
	ret, err := strconv.Atoi(strings.Split(string(status), " ")[1])
	if err != nil {
		fmt.Println(err)
	}
	return ret
}

func biggestProcess(memory []int) int {
	var max, ans int
	for i, mem := range memory {
		if mem > max {
			ans = i
			max = mem
		}
	}
	if max != 0 {
		return ans
	}else{
		return -1
	}
}

func killProcess(pid string) {
	pidInt, err := strconv.Atoi(pid[6:])
	if err != nil{
		fmt.Println(err)
	}
	proc, err := os.FindProcess(pidInt)
	if err != nil{
		fmt.Println(err)
	}
	proc.Kill()
}