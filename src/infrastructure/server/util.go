package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/cmd/config"

	"github.com/labstack/echo/v4"
)

func runWithGracefulShutdown(server *echo.Echo) {
	if server == nil {
		log.Fatal("server is nil")
	}

	httpServer := http.Server{
		Addr:    config.Config.Server.Port,
		Handler: server,
	}

	go func() {
		log.Println("server started")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen and serve error :", err)
		}
	}()

	q := make(chan os.Signal, 1)
	signal.Notify(q, os.Interrupt)

	<-q
	log.Println("server shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown error :", err)
	}

	log.Println("server exited properly")
}
