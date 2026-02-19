package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()

	// definir rutas que estaran presentes en handler.go
	// ejemplo "POST /shorten" o "GET /{code}"

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 15 * time.Second,
	}

	// iniciar servidor en una goroutine

	go func() {
		fmt.Println("Servidor corriendo en puerto 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error del servidor: %v \n", err)
		}
	}()

	// canal para escuchar al sistema operativo
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	fmt.Println("\nSeñal apagado recibida. Cerrando conexiones...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Servidor forzado a apagarse: %v", err)
	}


	// cerrar conexion DB

	fmt.Println("Servidor apagado correctamente.")

}