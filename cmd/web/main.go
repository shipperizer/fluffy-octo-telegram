package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"

	"github.com/shipperizer/fluffy-octo-telegram/pkg/routing"
)

type EnvSpec struct {
	HTTPPort    string `envconfig:"http_port"`
	JWKPubPath  string `envconfig:"jwk_pub_path"`
	JWKPrivPath string `envconfig:"jwk_priv_path"`
}

func main() {
	var specs EnvSpec
	err := envconfig.Process("", &specs)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.SetLevel(log.DebugLevel)

	apis, err := routing.NewRouter(specs.JWKPubPath, specs.JWKPrivPath)

	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%v", specs.HTTPPort),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      apis, // Pass our instance of gorilla/mux in.
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Debug("Shutting down")
	os.Exit(0)
}
