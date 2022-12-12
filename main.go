package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

func main() {
	addr := flag.String("addr", "0.0.0.0:8080", "Server listening addr")
	level := flag.String("level", "WARNING", "Server log level")
	flag.Parse()

	logLevel, err := logging.LogLevel(strings.ToUpper(*level))
	if err != nil {
		panic(err.Error())
	}
	leveledBackend.SetLevel(logLevel, log.Module)
	log.SetBackend(leveledBackend)

	gin.SetMode(gin.ReleaseMode)
	gin.Logger()
	gin.ForceConsoleColor()
	router := gin.Default()
	{
		router.GET("/health", HealthHandler)
		router.GET("/hello", HelloHandler)
		router.GET("/info", InfoHandler)
		router.GET("/", WebSocketHandler)
		router.POST("/subscribe", SubscribeHandler)
	}
	srv := &http.Server{Addr: *addr, Handler: router}
	go func() {
		var err error
		log.Info("Server Listen On:", "http://"+srv.Addr)
		err = srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server shutdown error: %s\n", err)
	} else {
		log.Info("Server exiting.")
	}
}
