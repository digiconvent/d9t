package api_engine

import (
	"crypto/tls"
	"net/http"

	"github.com/digiconvent/d9t/api/context"
	"github.com/digiconvent/d9t/meta/environment"
)

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

type Op struct {
	M Method
	H HandlerFunc
}

func Get(h HandlerFunc) Op    { return Op{GET, h} }
func Post(h HandlerFunc) Op   { return Op{POST, h} }
func Put(h HandlerFunc) Op    { return Op{PUT, h} }
func Delete(h HandlerFunc) Op { return Op{DELETE, h} }

type RouteTable map[string]map[string]map[string]Op

type HandlerFunc func(ctx *context.Context)

func SetupServer() *http.Server {
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
		Addr:      ":443",
		Handler:   SetupHandler(),
		TLSConfig: tlsConfig,
	}
	return srv

}

func SetupHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", ApiHandler)
	return mux
}
