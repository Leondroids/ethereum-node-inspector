package main

import (
	"log"
	"fmt"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
	"github.com/Leondroids/gox"
	"github.com/envion/golib"
	"github.com/Leondroids/ethereum-node-inspector/app"
)

func main() {
	context, err := app.InitApp()
	if err != nil {
		panic(err)
	}
	log.Printf("Config: %+v\n", context.Config)

	startServer(context)
}

func startServer(context *app.Context) {

	log.Println("Starting rpc server on port: ", context.Config.Port)

	handler := cors.Default().Handler(router(context))

	fmt.Println(http.ListenAndServe(context.Config.Port, handler))
}

func router(context *app.Context) *mux.Router {
	commonHandlers := alice.New(golib.LoggingHandler)

	// main router
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/healthcheck", commonHandlers.ThenFunc(gox.HealthcheckHandler)).Methods("GET")

	// node
	nodeHandler := app.NewNodeHandler(context)
	nodeRouter := router.PathPrefix("/node").Subrouter().StrictSlash(true)
	nodeRouter.Handle("/status", commonHandlers.ThenFunc(nodeHandler.Status)).Methods("GET")
	nodeRouter.Handle("/info", commonHandlers.ThenFunc(nodeHandler.Info)).Methods("GET")
	nodeRouter.Handle("/accounts", commonHandlers.ThenFunc(nodeHandler.Accounts)).Methods("GET")

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/"))))

	return router
}
