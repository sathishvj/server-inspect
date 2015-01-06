package main

import (
	"fmt"
	"time"

	"github.com/cloudfoundry/gosigar"
)

type MemInfo struct {
	Total, Used, Free uint64
	At                int64
}

var memDuration = 1 //seconds
var memMaxCount = 200
var memHistory []MemInfo

func memInfoSrvc(tickerDuration int) {
	fmt.Println("sysinfo.go:memInfoSrvc(): Starting memory ticker.")
	ticker := time.NewTicker(time.Duration(tickerDuration) * time.Second)
	go func() {
		for range ticker.C {
			mi := getMemInfo()
			// fmt.Println(mi)
			memHistory = append(memHistory, mi)
			if len(memHistory) > memMaxCount {
				memHistory = memHistory[1:]
			}
		}
	}()
}

func getMemInfo() MemInfo {
	mem := sigar.Mem{}
	mem.Get()
	return MemInfo{mem.Total, mem.Used, mem.Free, time.Now().UnixNano() / 1e6}
}

/*
type CpuInfo struct {
	At int64 //time in ms
}

var cpuDuration = 1 //seconds
var cpuMaxCount = 200
var cpuHistory []CpuInfo

func cpuInfoSrvc(tickerDuration int) {
	fmt.Println("sysinfo.go:cpuInfoSrvc(): Starting cpuory ticker.")
	ticker := time.NewTicker(time.Duration(tickerDuration) * time.Second)
	go func() {
		for range ticker.C {
			ci := getCpuInfo()
			// fmt.Println(mi)
			cpuHistory = append(cpuHistory, ci)
			if len(cpuHistory) > cpuMaxCount {
				cpuHistory = cpuHistory[1:]
			}
		}
	}()
}

func getCpuInfo() []CpuInfo {
	return nil
}
*/
