package main

import (
	"time"
	"log"
	"flag"
	"math/rand"
	"os"
	"os/signal"
	//"syscall"
)

func main() {
	// load command line arguments
	name := flag.String("name", "world", "name to print")
	flag.Parse()
	log.Printf("Starting sleepservice for %s", *name)
	// setup signal catching
	sigs := make(chan os.Signal, 1)
	// catch all signals since not explicitly listing
	signal.Notify(sigs)
	//signal.Notify(sigs,syscall.SIGQUIT)
	// method invoked upon seeing signal
	go func() {
		s := <-sigs
		log.Printf("RECEIVED SIGNAL: %s", s)
		AppCleanup()
		os.Exit(1)
	}()

	// infinite print loop
	for {
		log.Printf("hello %s", *name)
		// wait random number of milliseconds
		Nsecs := rand.Intn(3000)
		log.Printf("About to sleep %dms before looping again", Nsecs)
		time.Sleep(time.Millisecond * time.Duration(Nsecs))
	}
}
func AppCleanup() {
	log.Println("CLEANUP APP BEFORE EXIT!!!")
}
