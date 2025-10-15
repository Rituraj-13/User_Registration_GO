package app

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Rituraj-13/userReg/backend/internals/api"
	"github.com/Rituraj-13/userReg/backend/internals/store"
	"github.com/Rituraj-13/userReg/backend/migrations"
	"github.com/joho/godotenv"
)

type Application struct {
	Logger      *log.Logger
	DB          *sql.DB
	UserHandler *api.UserHandler
}

func NewApplication() (*Application, error) {

	godotenv.Load()

	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFs(pgDB, ".", migrations.FS)
	if err != nil {
		panic(err)
	}

	userStore := store.NewPostgresUserStore(pgDB)
	userHandler := api.NewUserHandler(userStore)

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	return &Application{
		Logger: logger,
		UserHandler: userHandler,
		DB: pgDB,
	}, nil
}

func (ap *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Everything is running fine and smooth !"))
	w.WriteHeader(http.StatusOK)
}
