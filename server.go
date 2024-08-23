package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jacky-htg/brilian/pkg/database"
	"github.com/jacky-htg/brilian/route"
)

func main() {

	logErr := log.New(os.Stdout, "ERROR : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	//info := log.New(os.Stdout, "INFO : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	db, err := database.OpenDB()
	if err != nil {
		log.Fatalf("error: connecting to db: %s", err)
	}
	defer db.Close()

	log.Fatal(http.ListenAndServe(":9000", route.InitRoute(db, logErr)))
	/*
		// parameter server
		server := http.Server{
			Addr:         "0.0.0.0:9000",
			Handler:      http.HandlerFunc(user.Get),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		}

		serverErrors := make(chan error, 1)
		// mulai listening server
		go func() {
			info.Println("server listening on", server.Addr)
			serverErrors <- server.ListenAndServe()
		}()

		// Membuat channel untuk mendengarkan sinyal interupsi/terminate dari OS.
		// Menggunakan channel buffered karena paket signal membutuhkannya.
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

		// Mengontrol penerimaan data dari channel,
		// jika ada error saat listenAndServe server maupun ada sinyal shutdown yang diterima
		select {
		case err := <-serverErrors:
			log.Fatalf("error: listening and serving: %s", err)

		case <-shutdown:
			info.Println("caught signal, shutting down")

			// Jika ada shutdown, meminta tambahan waktu 5 detik untuk menyelesaikan proses yang sedang berjalan.
			const timeout = 5 * time.Second
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				log.Printf("error: gracefully shutting down server: %s", err)
				if err := server.Close(); err != nil {
					log.Printf("error: closing server: %s", err)
				}
			}
		}

		info.Println("done")
	*/
}
