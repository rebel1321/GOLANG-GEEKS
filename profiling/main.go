package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // Import pprof for profiling via HTTP
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:8071", nil))
	}()
	

f, err := os.Create("cpu.prof")
if err != nil {
	log.Fatal(err)
}
defer f.Close()

if err := pprof.StartCPUProfile(f); err != nil {
	log.Fatal(err)
}
defer pprof.StopCPUProfile()

  for i := 0; i < 30; i++ {
		time.Sleep(1 * time.Second)
		simulateWorkload()	
	}
}

func simulateWorkload() {
	for i := 0; i < 1e7; i++ {
		_ = i * i
	}
}
