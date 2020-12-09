package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	app := http.NewServeMux()
	app.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello App"))
	})
	
	group, ctx := errgroup.WithContext(context.Background())
	group.Go(func() error {
		return server(ctx, ":8000", app, stop)
	})
	
	group.Go(func() error {
		return server(ctx, ":9000", http.DefaultServeMux, stop)
	})
	
	if err := group.Wait(); err != nil {
		fmt.Println(err)
		close(stop)
	}
	
}

func server(ctx context.Context, addr string, handler http.Handler, stop <-chan os.Signal) error {
	server := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	
	go func() {
		<-stop
		server.Shutdown(ctx)
	}()
	return server.ListenAndServe()
}
