package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	killSleep time.Duration
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func sigHandler(c chan os.Signal) {
	s := <-c
	log.Println("Got signal:", s)
	log.Printf("Sleeping for %s\n", killSleep)
	time.Sleep(killSleep)
	log.Println("Sleep Complete")
	os.Exit(0)
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)

	ks := os.Getenv("KILL_SLEEP")
	if ks == "" {
		killSleep = time.Duration(10 * time.Second)
	} else {
		secs, convErr := strconv.Atoi(ks)
		if convErr != nil {
			panic(convErr)
		}
		killSleep = time.Duration(time.Duration(int64(secs)) * time.Second)
	}

	go sigHandler(c)
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
