package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"os"
	"pkg/handler"
	"pkg/repository"
	"pkg/repository/db/postgres"
	"pkg/server"
	"pkg/service"
)

func main() {
	router := mux.NewRouter()

	db, err := postgres.New(postgres.Config{
		Host:     os.Getenv("DBHost"),
		Port:     os.Getenv("DBPort"),
		UserName: os.Getenv("DBUser"),
		Password: os.Getenv("DBPass"),
		DBName:   os.Getenv("DBName"),
		SSLMode:  os.Getenv("DBMode"),
	})
	if err != nil {
		log.Fatal(err)
	}

	r := repository.New(db)
	serv := service.NewService(r)
	h := handler.New(r, serv)
	h.Routes(router)

	s := server.New("8080", router)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
