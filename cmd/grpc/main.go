package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"

	"github.com/shipperizer/fluffy-octo-telegram/pkg/auth"
	"github.com/shipperizer/fluffy-octo-telegram/pkg/monitor"
)

type EnvSpec struct {
	HTTPPort string `envconfig:"http_port"`
	GRPCPort string `envconfig:"grpc_port"`
	CertPath string `envconfig:"cert_path"`
	CAPath   string `envconfig:"ca_path"`
	KeyPath  string `envconfig:"key_path"`
}

func main() {
	var specs EnvSpec
	err := envconfig.Process("", &specs)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.SetLevel(log.DebugLevel)

	// AWS NLB TLS termination.
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", specs.GRPCPort))

	if err != nil {
		log.Fatal("Error: ", err)
		return
	}

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%v", specs.HTTPPort),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      monitor.MetricsServer(), // Pass our instance of gorilla/mux in.
	}

	api := auth.NewGRPC(
		// &auth.TLS{
		// 	CAPath:   specs.CAPath,
		// 	CertPath: specs.CertPath,
		// 	KeyPath:  specs.KeyPath,
		// },
		nil
	)

	go func() {
		log.Infof("Starting GRPC Server up on %v", specs.GRPCPort)
		if err := api.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		log.Infof("Starting up on %v", specs.HTTPPort)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	// Block until we receive our signal.
	<-c

	api.GracefulStop()

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Info("Shutting down")
	os.Exit(0)
}
