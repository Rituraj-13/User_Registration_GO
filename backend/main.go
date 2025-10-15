package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/Rituraj-13/userReg/backend/internals/app"
	"github.com/Rituraj-13/userReg/backend/internals/routes"
)

func main() {
	app, err := app.NewApplication();
	if err != nil{
		panic(err)
	}

	var port int
	flag.IntVar(&port, "port", 8080, "Backend server port")
	flag.Parse()

	defer app.DB.Close()

	app.Logger.Printf("Server running on port : %d", port)
	routes := routes.SetupRoutes(app)

	server := &http.Server{
		Addr: fmt.Sprintf(":%d",port),
		Handler: routes,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout: time.Minute,
	}

	err = server.ListenAndServe()
	if err != nil{
		app.Logger.Fatal(err)
	}
}