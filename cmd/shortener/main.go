package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/BeInBloom/anima-sol/internal/app"
	defaultmuxbuilder "github.com/BeInBloom/anima-sol/internal/router_builder/default_mux_builder"
	"github.com/BeInBloom/anima-sol/models"
)

func main() {
	builder := defaultmuxbuilder.New()

	a := app.New(models.ServerDeps{
		Scheme:     "http",
		Host:       "localhost",
		Port:       8000,
		MuxBuilder: builder,
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println("server is starting...")

		if err := a.Run(); err != nil {
			if !errors.Is(err, app.ErrAppClosed) {
				fmt.Printf("error: %s\n", err)
				return
			}

			return
		}
	}()

	<-stop

	fmt.Println("its over...")

	if err := a.Close(); err != nil {
		fmt.Printf("error: %s", err)
	}
}
