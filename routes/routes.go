package routes

import (
	"log"
	"net/http"

	"github.com/MarcosRoch4/api-go-sqlite/controllers"
	"github.com/MarcosRoch4/api-go-sqlite/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)

	r.HandleFunc("/api/posto", controllers.TodosPosto).Methods("Get")
	r.HandleFunc("/api/posto", controllers.NovoPosto).Methods("Post")
	r.HandleFunc("/api/posto/{id}", controllers.RetornaUmPosto).Methods("Get")
	r.HandleFunc("/api/posto/{id}", controllers.DeletaPosto).Methods("Delete")
	r.HandleFunc("/api/posto/{id}", controllers.EditaPosto).Methods("Put")

	r.HandleFunc("/api/combustivel", controllers.TodosCombustiveis).Methods("Get")
	r.HandleFunc("/api/combustivel", controllers.NovoCombustivel).Methods("Post")
	r.HandleFunc("/api/combustivel/{id}", controllers.RetornaUmCombustivel).Methods("Get")
	r.HandleFunc("/api/combustivel/{id}", controllers.DeletaCombustivel).Methods("Delete")
	r.HandleFunc("/api/combustivel/{id}", controllers.EditaCombustivel).Methods("Put")

	r.HandleFunc("/api/valor", controllers.TodosValores).Methods("Get")
	r.HandleFunc("/api/valor", controllers.NovoValor).Methods("Post")
	r.HandleFunc("/api/valor/{id}", controllers.RetornaUmValor).Methods("Get")
	r.HandleFunc("/api/valor/{id}", controllers.DeletaValor).Methods("Delete")
	r.HandleFunc("/api/valor/{id}", controllers.EditaValor).Methods("Put")

	r.HandleFunc("/api/categoria", controllers.TodasCategorias).Methods("Get")
	r.HandleFunc("/api/categoria", controllers.NovaCategoria).Methods("Post")
	r.HandleFunc("/api/categoria/{id}", controllers.RetornaUmaCategoria).Methods("Get")
	r.HandleFunc("/api/categoria/{id}", controllers.DeletaCategoria).Methods("Delete")
	r.HandleFunc("/api/categoria/{id}", controllers.EditaCategoria).Methods("Put")

	//r.HandleFunc("/api/ordenados", controllers.PostoPorCategoria).Methods("Get")
	r.HandleFunc("/api/ordenados/categoria", controllers.PostoPorCategoria).Methods("Get")
	r.HandleFunc("/api/ordenados/combustivel/{id}", controllers.PostoPorCombustivel).Methods("Get")

	r.HandleFunc("/api/ordenados/resultado", controllers.Resultado).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
