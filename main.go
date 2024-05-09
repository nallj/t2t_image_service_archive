package main

import (
	"github.com/nallj/t2t_image_service/config"

	"errors"
	"log"
	"net"
	"net/http"
	"sync"
)

const (
	HEALTHZ_ROUTE string = "/healthz"
	IMAGE_ROUTE   string = "/image"
	JWT_ROUTE     string = "/jwt"
)

func main() {
	log.Print("Starting image service...")

	cfg := config.NewConfig()
	cfg.InitConfig()

	port := cfg.Port
	// wat do?

	httpServer := &http.Server{
		Handler: http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			log.Printf("Recieved request: '%s'", req.URL.Path)

			switch req.URL.Path {
			case HEALTHZ_ROUTE:
				handleHealthz(writer)
				break
			case IMAGE_ROUTE:
				handleImage(cfg, writer, req)
				break
			case JWT_ROUTE:
				handleJwt(cfg, writer, req)
				break
			default:
				writer.WriteHeader(404)
			}
		}),
	}

	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	log.Print("Wait group add...")
	go func() {
		log.Print("Go routing running...")
		defer waitGroup.Done()

		addr := ":" + port
		listener, listenErr := net.Listen("tcp", addr)
		if listenErr != nil {
			log.Fatalf("Failed to listen: %v", listenErr)
		}
		defer listener.Close()
		defer httpServer.Close()

		host := "http://localhost:" + port + HEALTHZ_ROUTE
		log.Printf("Listening on %s", host)
		serverErr := httpServer.Serve(listener)
		if !errors.Is(serverErr, http.ErrServerClosed) {
			log.Fatalf("Failed to listen and serve: %v", serverErr)
		}
		log.Print("Go routing finished...")
	}()

	// Only finish when no wait groups remain.
	waitGroup.Wait()

	log.Print("Go routing go...")
}
