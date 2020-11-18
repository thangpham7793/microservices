// Setting Up go modules: https://blog.francium.tech/go-modules-go-project-set-up-without-gopath-1ae601a4e868

package main

import (
	"context"
	"lesson2/handlers"
	"lesson2/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	envs := utils.GetEnvs(".env")

	// customise dependency (anything that implements the interface, in this case io.Writer)
	l := log.New(os.Stdout, "product-api\t", log.LstdFlags)

	hh := handlers.NewHello(l)
	gb := handlers.NewGoodbye(l)

	// like the controller/router in MVC
	sm := http.NewServeMux()

	sm.Handle("/", hh)
	sm.Handle("/goodbye", gb)

	s := &http.Server{
		Addr:         envs["Addr"],
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// so as not to block shutdown setup
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// block until a signal is received
	sigChan := make(chan os.Signal)

	// send signal to channel
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// like callback (when received, then do work)
	sig := <-sigChan
	l.Printf("Received signal %s, graceful shutdown\n", sig)

	// TODO: need to read up on context package
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
