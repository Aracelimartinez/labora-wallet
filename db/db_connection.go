package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Driver de conexión con Postgres
)

type DbData struct {
	Host        string
	Port        string
	DbName      string
	RolName     string
	RolPassword string
}

var DbConn *sql.DB

// Abre la conexión con la base de datos
func EstablishDbConnection() error {

	dbData, err := LoadEnv()

	if err != nil {
		log.Fatal(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbData.Host, dbData.Port, dbData.RolName, dbData.RolPassword, dbData.DbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	DbConn = db

	fmt.Println("Conexión exitosa a la base de datos:", DbConn)

	if err = DbConn.Ping(); err != nil {
		DbConn.Close()
		log.Fatal(err)
	}

	return nil
}

// Carga los datos del archivo .env
func LoadEnv() (DbData, error) {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
		return DbData{}, err
	}

	return DbData{
		Host:        os.Getenv("HOST"),
		Port:        os.Getenv("PORT"),
		DbName:      os.Getenv("DB_NAME"),
		RolName:     os.Getenv("ROL_NAME"),
		RolPassword: os.Getenv("ROL_PASSWORD"),
	}, nil
}

// Inicia el servidor
func StartServer(port string, router http.Handler) error {
	server := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Iniciando el servidor en el puerto %s..\n", port)

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("error al iniciar el servidor: %v", err)
	}

	return nil
}
