package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/brandoyts/go-clean/infrastructure/database/mongodb"
	"github.com/brandoyts/go-clean/infrastructure/rest/routes"
	"github.com/brandoyts/go-clean/internal/controller"
	"github.com/brandoyts/go-clean/internal/repository/mongoRepository"
	"github.com/brandoyts/go-clean/internal/service"
)

type Controller struct {
	UserController *controller.UserController
	AuthController *controller.AuthController
}

type Server struct {
	Port       string
	Database   *mongodb.Mongodb
	Controller *Controller
}

func NewServer(port string) *Server {
	return &Server{
		Port: port,
	}
}

func (s *Server) registerDependencies() {
	db, err := mongodb.New("go-clean", "mongodb://gomongouser:secretpassword@localhost:27017/?authSource=admin")
	if err != nil {
		panic(err)
	}

	log.Println("✅ successfully connected to mongodb")

	userRepository := mongoRepository.NewUserMongoRepository(db.Db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	authService := service.NewAuthService(userRepository)
	authController := controller.NewAuthController(*authService)

	s.Database = db

	s.Controller = &Controller{
		UserController: userController,
		AuthController: authController,
	}
}

func (s *Server) Start() {
	s.registerDependencies()

	router := routes.Initialize(
		s.Controller.UserController,
		s.Controller.AuthController,
	)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", s.Port),
		Handler: router,
	}

	go func() {
		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := httpServer.Shutdown(shutdownCtx)
	if err != nil {
		log.Println("Server Shutdown:", err)
	}

	<-shutdownCtx.Done()
	log.Println("timeout of 3 seconds.")
	log.Println("Server exiting")
}
