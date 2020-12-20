package main

import (
	"context"
	"goTraining/Week04/internal/setting"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	"golang.org/x/sync/errgroup"
	
	"github.com/gorilla/mux"
	
	_ "time/tzdata"
)

func main() {
	stopCh := setUpSignalHandler()
	group, _ := errgroup.WithContext(context.Background())
	h := InitHandler(setting.DB)
	
	r := mux.NewRouter()
	r.HandleFunc("/v1/save/{name}", h.Save).Methods(http.MethodPost)
	r.HandleFunc("/v1/get/{name}", h.Get).Methods(http.MethodGet)
	
	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	
	group.Go(func() error {
		return server.ListenAndServe()
	})
	
	group.Go(func() error {
		<-stopCh
		return server.Shutdown(context.Background())
	})
	
	if err := group.Wait(); err != nil {
		log.Fatalf("errgroup: %v", err)
	}
}

func setUpSignalHandler() <-chan struct{} {
	stop := make(chan struct{})
	
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	
	go func() {
		<-c
		close(stop)
		<-c
	}()
	
	return stop
}
