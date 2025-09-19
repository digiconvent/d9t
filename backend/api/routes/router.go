package routes

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/digiconvent/d9t/meta/environment"
	"github.com/digiconvent/d9t/meta/services"
)

func InitRouter(s *services.Services) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hjallo djallo")
	})

	certPEM := []byte(environment.Env.TlsCertificate)
	keyPEM := []byte(environment.Env.TlsPrivateKey)

	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	srv := &http.Server{
		Addr:      ":80",
		Handler:   mux,
		TLSConfig: tlsConfig,
	}
	return srv
}
