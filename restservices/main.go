// Setting Up go modules: https://blog.francium.tech/go-modules-go-project-set-up-without-gopath-1ae601a4e868

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"restservices/handlers"
	"restservices/utils"
	"time"
)

func main() {

	envs := utils.GetEnvs(".env")

	// customise dependency (anything that implements the interface, in this case io.Writer)
	l := log.New(os.Stdout, "product-api\t", log.LstdFlags)

	//dependency injections allow for reusable dependencies (logger for example)
	// hh := handlers.NewHello(l)
	// gb := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)

	// like the controller/router in MVC
	sm := http.NewServeMux()

	// sm.Handle("/", hh)
	// sm.Handle("/goodbye", gb)
	sm.Handle("/products/", ph)

	s := &http.Server{
		Addr:         envs["Addr"],
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 10 * time.Second,
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

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
