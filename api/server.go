package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	userHttpDelivery "github.com/Lorengamboa/whitelabel-travel/user/delivery/http"
	userRepository "github.com/Lorengamboa/whitelabel-travel/user/repository/pg"
	userUsecase "github.com/Lorengamboa/whitelabel-travel/user/usecase"
)

const port = "8080"

func main() {
	db, err := sql.Open("postgres", os.Getenv("POSTGRESQL_URL"))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	mux := http.NewServeMux()

	timeoutContext := time.Duration(10) * time.Second

	// User
	userRepository := userRepository.NewPgUserRepository(db)
	userUsecase := userUsecase.NewUserUsecase(userRepository, timeoutContext)

	// Handlers
	userHttpDelivery.NewUserHandler(mux, userUsecase)

	slog.Info(fmt.Sprintf("server listening at http://localhost:%s", port))
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
