package server

import (
	"log"
	"net/http"

	"github.com/LiFeAiR/users-crud-ai/internal/handlers"
	"github.com/LiFeAiR/users-crud-ai/internal/repository"
)

// Server представляет HTTP сервер
type Server struct {
	port string
	db   *repository.DB
}

// NewServer создает новый экземпляр сервера
func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

// DB возвращает подключение к базе данных
func (s *Server) DB() *repository.DB {
	return s.db
}

// Start запускает HTTP сервер
func (s *Server) Start(connStr string) error {
	// Подключаемся к базе данных
	db, err := repository.NewDB(connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	s.db = db

	baseHandler := handlers.NewBaseHandler(repository.NewUserRepository(db))

	//// Определяем обработчик для корневого пути
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetRootHandler(w, r)
	})

	// Определяем обработчик для эндпоинта /api/users
	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			baseHandler.GetUsersHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Определяем обработчик для эндпоинта /api/user с несколькими методами
	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Определяем обработчик для GET /api/user/id=int
			baseHandler.GetUser(w, r)
		case http.MethodPost:
			baseHandler.CreateUser(w, r)
		case http.MethodPut:
			baseHandler.UpdateUser(w, r)
		case http.MethodDelete:
			baseHandler.DeleteUser(w, r)
		}
	})

	// Запускаем HTTP сервер
	log.Printf("Starting HTTP server on port %s...\n", s.port)
	log.Fatal(http.ListenAndServe(":"+s.port, nil))

	return nil
}
