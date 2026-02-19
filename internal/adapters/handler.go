package adapters

import (
	"encoding/json"
	"net/http"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url,omitempty"`
	Error    string `json:"error,omitempty"`
}

// handler para GET /home. Devuelve pagina principal.
func HandleHome(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type",)
	w.WriteHeader(http.StatusOK)
	//w.Write()
}



// handler para POST /shorten. Recibe por JSON la url larga y devuelve la corta.
func HandleShorten(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendJSONError(w, "FORMATO JSON INVALIDO", http.StatusBadRequest)
		return
	}

	// aqui se llaman a validaciones internas de internal/core

	// aqui se llama al servicio de para generar y guardar el codigo
	code := "123456" // ejmplo

	response := ShortenResponse{
		ShortURL: "http://localhost:8080/" + code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(response)
}



// handler para GET /{code}. Redirije a la url original segun el codigo.
func HandlerRedirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	if code == "" {
		http.Error(w, "CODIGO NO PROPORCIONADO", http.StatusBadRequest)
		return
	}

	// aqui se buscaria la URL original en la DB
	// originalURL, err := core.GetURL(code)
	originalURL := "https://example.com" // ejemplo

	// usamos Found (302) en lugar de MovedPermanently (301) para que el navegador no cachee la redireccion y poder contar cada visita dee un mismo usuario
	http.Redirect(w, r, originalURL, http.StatusFound)

}



// handler para GET /stats/{code}. Devuelve las estadisticas de la url acortada. De momento son visitas y ultima visita (fecha).
func HandleStats(w http.ResponseWriter, r *http.Request) {

}



// funcion auxiliar para modularizar el codigo de enviar error
func sendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ShortenResponse{Error: message})
}