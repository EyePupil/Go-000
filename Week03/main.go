package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"os/signal"
	"routehttp/nonstd/sync/errgroup"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Second * 10)
		writer.WriteHeader(200)
		_, _ = writer.Write([]byte("OK"))
	})

	server := &http.Server{Addr:":8080", Handler:mux}
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		fmt.Println("start listen on port", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			return errors.Wrap(err, "listen server")
		}
		return nil
	})

	g.Go(func() error {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <- sc:
			fmt.Println("receive signal")
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				fmt.Println("exit...")
				ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
				defer cancel()
				if err := server.Shutdown(ctx); err != nil {
					return errors.Wrap(err, "shutdown server")
				}
				return nil
			case syscall.SIGHUP:
				fmt.Println("restart server")
			default:
				fmt.Println("unknown signal")
			}
		case <- ctx.Done():
			fmt.Println("exit listen")
			return nil
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("%+v", err)
	}
}