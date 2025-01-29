package routes

import (
	"net/http"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/internal/handlers"
)

type Router struct {
	routes  map[string]handlers.Handler
	baseDir string
}

func NewRouter(baseDir string) *Router {
	return &Router{
		routes: map[string]handlers.Handler{
			"/": &handlers.HomeHandler{},
		},
		baseDir: baseDir,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// Manejo de rutas dinámicas "/file/{filename}"
	if strings.HasPrefix(req.URL.Path, "/file/") {
		filename := strings.TrimPrefix(req.URL.Path, "/file/")
		fileHandler := handlers.NewFileHandler(r.baseDir, filename)
		fileHandler.ServeHTTP(w, req)
		return
	}

	// Manejo de rutas dinámicas "/echo/{mensaje}"
	if strings.HasPrefix(req.URL.Path, "/echo/") {
		message := strings.TrimPrefix(req.URL.Path, "/echo/")
		echoHandler := handlers.NewEchoHandler(message)
		echoHandler.ServeHTTP(w, req)
		return
	}

	// Rutas estáticas
	if handler, exists := r.routes[req.URL.Path]; exists {
		handler.ServeHTTP(w, req)
	} else {
		(&handlers.NotFoundHandler{}).ServeHTTP(w, req)
	}
}
