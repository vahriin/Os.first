package main

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"io/ioutil"
	"errors"
)

func main() {
	PIDs, err := getAllProcessPids()
	if err != nil {
		panic(err)
	}

	memPID := getUsedMemory(PIDs)
	posBigProc, err := biggestProcess(memPID)
	if err != nil {
		panic(err)
	}

	killProcess(PIDs[posBigProc])
}

func getAllProcessPids() ([]string, error) {
	return filepath.Glob("/proc/[0-9]*") //get all files in /proc

}

func getUsedMemory(pids []string) ([]int) {
	memOfProc := make([]int, len(pids))
	for i, pid := range pids{
		status, err := ioutil.ReadFile(pid+"/statm")
		if err != nil {
			panic(err)
		}

		memOfProc[i], err = strconv.Atoi(strings.Split(string(status), " ")[1])
		if err != nil {
			panic(err)
		}
	}
	return memOfProc //с производительностью все ок, слайсы передаются по указателю
}

func biggestProcess(memory []int) (int, error) {
	var max, ans int
	if len(memory) == 0 {
		return 0, errors.New("No values in memory array")
	}
	for i, mem := range memory {
		if mem > max {
			ans = i
			max = mem
		}
	}
	if max == 0 {
		return 0, errors.New("Max values is null")
	} else {
		return ans, nil
	}
}

func killProcess(pid string) {
	pidInt, err := strconv.Atoi(pid[6:])
	if err != nil{
		panic(err)
	}
	proc, err := os.FindProcess(pidInt)
	if err != nil{
		panic(err)
	}
	proc.Kill()
}