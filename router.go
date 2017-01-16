package main

import (
	"net/http"
	"crypto/tls"
)


func setupRouting(overrideHttpPort string, cfg map[string]string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ConfigureHandleAction(cfg))

	https_port, use_https := cfg["https_port"]

	if use_https {
		go func() {
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
			err := srv.ListenAndServeTLS(cfg["tls_crt_file"], cfg["tls_key_file"])

			if err != nil {
				panic(err)
			}
		}()
	}

	port := cfg["http_port"]

	if overrideHttpPort != "" {
		port = overrideHttpPort
	}

	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		panic(err)
	}
}