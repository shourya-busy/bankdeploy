package database

import (
	"crypto/tls"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
)

var Db *pg.DB

func Connect() {

	addr := os.Getenv("DB_ADDR")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	opts := &pg.Options{
		User : user,
		Password: password,
		Addr : addr,
		Database: database,
        TLSConfig: &tls.Config{
            // Configure TLS version and cipher suite
            MinVersion: tls.VersionTLS13,
            CipherSuites: []uint16{
                tls.TLS_AES_128_GCM_SHA256,
            },
			PreferServerCipherSuites: true,
			InsecureSkipVerify:       true,
        },
	}

	Db = pg.Connect(opts)

	if Db == nil {
		log.Printf("Failed to connect to database.\n")
		os.Exit(100)
	}

}
