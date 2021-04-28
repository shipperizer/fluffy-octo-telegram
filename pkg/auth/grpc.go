package auth

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/google/martian/log"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

// InitGRPC initializes RPCs
func InitGRPC(tlsConfig *TLS) *grpc.Server {

	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	}

	if tlsConfig != nil {
		creds, err := NewServerCredentials(tlsConfig.CertPath, tlsConfig.KeyPath, tlsConfig.CAPath)

		if err != nil {
			log.Error("issue with tlsConfig for GRPC")
			log.Error("proceeding without tls for GRPC server")
		} else {
			opts.append(creds)
		}
	}

	// plugin prometheus monitoring
	server := grpc.NewServer(opts...)

	NewRPCv2().Register(server)
	NewRPCv3().Register(server)

	// TODO @shipperizer this increases cardinality of metric, use an env var for enabling it
	grpc_prometheus.EnableHandlingTimeHistogram()
	// init prometheus metrics provided by grpc_prometheus
	grpc_prometheus.Register(server)

	// Register reflection service on gRPC server
	reflection.Register(server)

	return server
}

// NewServerCredentials loads TLS transport credentials for the GRPC server.
func NewServerCredentials(certPath, keyPath, caPath string) (credentials.TransportCredentials, error) {
	srv, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	p := x509.NewCertPool()

	if caPath != "" {
		ca, err := ioutil.ReadFile(caPath) //nolint(gosec)
		if err != nil {
			return nil, err
		}

		p.AppendCertsFromPEM(ca)
	}

	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{srv},
		RootCAs:      p,
	}), nil
}

// NewGRPC creates a GRPC server
func NewGRPC(tlsConfig *TLS) *grpc.Server {
	return InitGRPC(tlsConfig)
}
