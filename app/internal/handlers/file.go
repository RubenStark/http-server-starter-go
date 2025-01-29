// internal/handlers/file.go
package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type FileHandler struct {
	BaseDir  string
	Filename string
}

// Constructor
func NewFileHandler(baseDir, filename string) *FileHandler {
	return &FileHandler{BaseDir: baseDir, Filename: filename}
}

func (h *FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGet(w, r)
	case http.MethodPost:
		h.handlePost(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// Manejo de GET: Leer y devolver el archivo
func (h *FileHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join(".", h.BaseDir, h.Filename)

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Archivo no encontrado", http.StatusNotFound)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, file)
}

// Manejo de POST: Guardar el contenido en un archivo
func (h *FileHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join(".", h.BaseDir, h.Filename)

	// Crear o sobrescribir el archivo
	file, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "No se pudo crear el archivo", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Copiar el contenido del body al archivo
	_, err = io.Copy(file, r.Body)
	if err != nil {
		http.Error(w, "Error al escribir en el archivo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Archivo %s guardado exitosamente", h.Filename)
}
