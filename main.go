package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"log"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available CPUs

	if len(os.Args) < 2 {
		fmt.Println("Usage: indexer file.tar.gz")
		os.Exit(1)
	}

	f, err := os.Create("cpu_profile.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()

	// Start the CPU profiler
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	filename := os.Args[1]
	err = readAndProcessTarGz(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
