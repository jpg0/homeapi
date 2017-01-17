package main

import (
	"net/http"
	"crypto/tls"
	"github.com/Sirupsen/logrus"
	"golang.org/x/crypto/acme/autocert"
)


func setupRouting(overrideHttpPort string, cfg *Configuration) {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", ConfigureHandleAction(cfg))

	https_port := cfg.HTTPSPort

	if https_port != "" {
		logrus.Infof("Serving SSL on port %v", https_port)

		serveSSLWithKeyfiles(https_port, mux, cfg)
	}

	port := cfg.HTTPPort

	if overrideHttpPort != "" {
		port = overrideHttpPort
	}

	logrus.Fatal(http.ListenAndServe(":" + port, nil))
}

func serveSSLWithAutocert(https_port string, mux *http.ServeMux, cfg map[string]string){
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(cfg["autocert_domain"]), //your domain here
		Cache:      autocert.DirCache("certs"), //folder for storing certificates
	}

	server := &http.Server{
		Addr: ":" + https_port,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
		Handler:      mux,
	}

	logrus.Fatal(server.ListenAndServeTLS("", "")) //key and cert are comming from Let's Encrypt
}

func serveSSLWithKeyfiles(https_port string, mux *http.ServeMux, cfg *Configuration){

	tlscfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         ":" + https_port,
		Handler:      mux,
		TLSConfig:    tlscfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	logrus.Fatal(srv.ListenAndServeTLS(cfg.Resolve(cfg.TLSCrtFile), cfg.Resolve(cfg.TLSKeyFile)))
}