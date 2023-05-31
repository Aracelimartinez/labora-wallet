package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DbData struct {
	Host        string
	Port        string
	DbName      string
	RolName     string
	RolPassword string
}

type Config struct {
	DbData       DbData
	TruoraAPIKey string
}

var DbConn *sql.DB
var GlobalConfig Config

// Open the conexion with the database
func EstablishDbConnection() error {

	config, err := LoadEnv()

	if err != nil {
		log.Fatal(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DbData.Host, config.DbData.Port, config.DbData.RolName, config.DbData.RolPassword, config.DbData.DbName)

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

// Load the data of the .env file
func LoadEnv() (Config, error) {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
		return Config{}, err
	}

	var config Config

	config = Config {
		DbData: DbData {
			Host:        os.Getenv("HOST"),
			Port:        os.Getenv("PORT"),
			DbName:      os.Getenv("DB_NAME"),
			RolName:     os.Getenv("ROL_NAME"),
			RolPassword: os.Getenv("ROL_PASSWORD"),
		},
		TruoraAPIKey: os.Getenv("TRUORA_API_KEY"),
	}

	GlobalConfig = config
	return config, nil
}

// Start the server
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
