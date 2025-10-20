package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

var db *sql.DB

func main() {
	var err error

	// Use internal Docker network port (5432) for Postgres
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),   // service name in Docker Compose
		getEnv("DB_PORT", "5433"),       // internal container port
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "password"),
		getEnv("DB_NAME", "mydb"),
	)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}
	log.Println("Connected to DB")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		password TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// HTTP handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/users", getUsers)
	mux.HandleFunc("/add", addUser)

	// Enable CORS only for frontend host (localhost:3000)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173","http://frontend:80"}, // change if frontend host changes
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT name, password FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.Name, &u.Password)
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO users (name, password) VALUES ($1, $2)", u.Name, u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User added successfully"))
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
